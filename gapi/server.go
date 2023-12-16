package gapi

import (
	"fmt"
	"github.com/kothar/go-backblaze"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/longIdt2502/pharmago_be/woker"
)

type ServerGRPC struct {
	pb.UnimplementedPharmagoServer
	config          utils.Config
	store           *db.Store
	tokenMaker      token.Maker
	taskDistributor woker.TaskDistributor
	b2Bucket        *backblaze.Bucket
}

func NewServer(config utils.Config, store *db.Store, taskDistributor woker.TaskDistributor, b2Bucket *backblaze.Bucket) (*ServerGRPC, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err)
	}

	server := &ServerGRPC{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
		b2Bucket:        b2Bucket,
	}

	return server, nil
}
