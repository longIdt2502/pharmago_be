package woker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendFcm = "task:send_fcm"

type PayloadSendFcm struct {
	To    string `json:"to"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (distributor *RedisTaskDistributor) DistributorTaskSendFcm(
	ctx context.Context,
	payload *PayloadSendFcm,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendFcm, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task:, %w", err)
	}
	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueue task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessorTaskSendFcm(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendFcm
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	err := processor.client.SendMessage(payload.To, payload.Title, payload.Body)
	if err != nil {
		log.Error().Str("channel", payload.To).Msg("can't send fcm message")
		return err
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("channel", payload.To).
		Msg("processor task")
	return nil
}
