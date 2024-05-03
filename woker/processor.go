package woker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/longIdt2502/pharmago_be/b2"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/firebase"
	"github.com/longIdt2502/pharmago_be/mail"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessorTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessorTaskSendOrderZns(ctx context.Context, task *asynq.Task) error
	ProcessorTaskSendFcm(ctx context.Context, task *asynq.Task) error
	ProcessorUploadImageVariant(ctx context.Context, task *asynq.Task) error
	ProcessorUploadImageProduct(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server   *asynq.Server
	store    db.Store
	mailer   mail.EmailSender
	client   *firebase.FCM
	b2Bucket *b2.B2Bucket
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender, client *firebase.FCM, b2Bucket *b2.B2Bucket) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().
					Err(err).
					Str("type", task.Type()).
					Bytes("payload", task.Payload()).
					Msg("process task failed")
			}),
			Logger: NewLogger(),
		},
	)

	return &RedisTaskProcessor{
		server:   server,
		store:    store,
		mailer:   mailer,
		client:   client,
		b2Bucket: b2Bucket,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessorTaskSendVerifyEmail)
	mux.HandleFunc(TaskSendOrderZns, processor.ProcessorTaskSendOrderZns)
	mux.HandleFunc(TaskSendFcm, processor.ProcessorTaskSendFcm)
	mux.HandleFunc(TaskUploadImageVariant, processor.ProcessorUploadImageVariant)
	mux.HandleFunc(TaskUploadImageProduct, processor.ProcessorUploadImageProduct)

	return processor.server.Start(mux)
}
