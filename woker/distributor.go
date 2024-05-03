package woker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
	DistributorTaskSendOrderZns(ctx context.Context, payload *PayloadZNS, opts ...asynq.Option) error
	DistributorTaskSendFcm(ctx context.Context, payload *PayloadSendFcm, opts ...asynq.Option) error
	DistributorUploadImageVariant(ctx context.Context, payload *PayloadUploadImageVariant, opts ...asynq.Option) error
	DistributorUploadImageProduct(ctx context.Context, payload *PayloadUploadImageProduct, opts ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
