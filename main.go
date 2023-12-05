package main

import (
	"context"
	"database/sql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	_ "github.com/longIdt2502/pharmago_be/docs/statik"
	"github.com/longIdt2502/pharmago_be/gapi"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open database:", err)
	}

	store := db.NewStore(conn)
	go runGatewayServer(config, &store)
	runServerGRPC(config, &store)

	//ginServer, err := api.NewServer(config, &store)
	//
	//err = ginServer.Start(config.HTTPServerAddress)
	//if err != nil {
	//	log.Fatal("cannot start server:", err)
	//}
}

func runServerGRPC(config utils.Config, store *db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPharmagoServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start GRPC server:", err)
	}
}

func runGatewayServer(config utils.Config, store *db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
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
		log.Fatal("cannot register handle server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	//fs := http.FileServer(http.Dir("./doc/swagger"))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik fs:", err)
	}
	swaggerHandle := http.StripPrefix("/swagger", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandle)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server")
	}
}
