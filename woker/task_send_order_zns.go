package woker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendOrderZns = "task:send_order_zns"
	UrlOa            = "https://core.wezolo.com/v1/zns/send/"
)

type PayloadZNS struct {
	OaID  string         `json:"oa_id"`
	Data  PayloadZNSData `json:"data"`
	Phone string         `json:"phone"`
	Mode  string         `json:"mode"`
	Type  string         `json:"type"`
}

type PayloadZNSData struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	OrderItems string `json:"order_items"`
	CreatedAt  string `json:"created_at"`
	Total      string `json:"total_price"`
	Phone      string `json:"phone"`
	Code       string `json:"order_code"`
}

func (distributor *RedisTaskDistributor) DistributorTaskSendOrderZns(ctx context.Context, payload *PayloadZNS, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendOrderZns, jsonPayload, opts...)
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

func (processor *RedisTaskProcessor) ProcessorTaskSendOrderZns(ctx context.Context, task *asynq.Task) error {
	var payload PayloadZNS
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	jsonStr, _ := json.Marshal(payload)

	client := &http.Client{}
	url := "https://core.wezolo.com/v1/zns/send/"
	reqHttp, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("Authorization", "Token 1e21c13f941d67507d9d1099150866b6759d9336")
	reqHttp.Header.Set("X-Application-Id", "4367014412197839377")
	reqHttp.Header.Set("X-Secret-Key", "6Sr1vfB1Jgh7ll7QkUmU")

	resp, err := client.Do(reqHttp)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Bytes("response", body).
		Str("phone", payload.Phone).
		Msg("processor task")

	return nil
}
