package main

import (
	"context"
	"database/sql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/kothar/go-backblaze"
	_ "github.com/lib/pq"
	"github.com/longIdt2502/pharmago_be/b2"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	_ "github.com/longIdt2502/pharmago_be/docs/statik"
	"github.com/longIdt2502/pharmago_be/gapi"
	config2 "github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/mail"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/longIdt2502/pharmago_be/woker"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open database")
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := woker.NewRedisTaskDistributor(redisOpt)

	b2Bucket, err := b2.NewB2Bucket(config.B2AccountId, config.B2ApplicationKey, config.B2KeyId, config.B2Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to b2")
	}

	go runTaskProcessor(config, redisOpt, store)
	go runGatewayServer(config, &store, taskDistributor, b2Bucket)
	runServerGRPC(config, &store, taskDistributor, b2Bucket)

	//ginServer, err := api.NewServer(config, &store)
	//
	//err = ginServer.Start(config.HTTPServerAddress)
	//if err != nil {
	//	log.Fatal("cannot start server:", err)
	//}
}

func runTaskProcessor(config utils.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := woker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runServerGRPC(config utils.Config, store *db.Store, taskDistributor woker.TaskDistributor, b2Bucket *backblaze.Bucket) {
	server, err := gapi.NewServer(config, store, taskDistributor, b2Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(config2.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterPharmagoServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start GRPC server")
	}
}

func runGatewayServer(config utils.Config, store *db.Store, taskDistributor woker.TaskDistributor, b2Bucket *backblaze.Bucket) {
	server, err := gapi.NewServer(config, store, taskDistributor, b2Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	option := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(option)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterPharmagoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handle server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	//fs := http.FileServer(http.Dir("./doc/swagger"))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}
	swaggerHandle := http.StripPrefix("/swagger", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandle)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := config2.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}
