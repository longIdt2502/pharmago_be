package gapi

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/woker"
)

func (server *ServerGRPC) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}
	account, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	payloadTx := db.CreateProductTxParams{
		CreateProductRequest: req,
		Account:              account,
		TokenPayload:         tokenPayload,
		UploadImageVariant: func(idVariant int32, image []byte) {
			payload := woker.PayloadUploadImageVariant{
				Image: image,
				Id:    idVariant,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(0),
				asynq.ProcessIn(1 * time.Second),
				asynq.Queue(woker.QueueCritical),
			}

			_ = server.taskDistributor.DistributorUploadImageVariant(ctx, &payload, opts...)
		},
	}

	productId, err := server.store.CreateProductTx(ctx, payloadTx)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Product.GetImage() {
		payload := woker.PayloadUploadImageProduct{
			Image: item,
			Id:    productId,
		}

		opts := []asynq.Option{
			asynq.MaxRetry(0),
			asynq.ProcessIn(1 * time.Second),
			asynq.Queue(woker.QueueCritical),
		}

		_ = server.taskDistributor.DistributorUploadImageProduct(ctx, &payload, opts...)
	}

	return &pb.CreateProductResponse{
		Message: "success",
		Code:    200,
		Details: productId,
	}, nil
}
