package gapi

import (
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/mail"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
)

type ServerGRPC struct {
	pb.UnimplementedPharmagoServer
	config     utils.Config
	store      *db.Store
	tokenMaker token.Maker
	sender     mail.EmailSender
}

func NewServer(config utils.Config, store *db.Store, sender mail.EmailSender) (*ServerGRPC, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err)
	}

	server := &ServerGRPC{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		sender:     sender,
	}

	return server, nil
}
