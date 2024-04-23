package woker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hibiken/asynq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/rs/zerolog/log"
)

const TaskSendFcm = "task:send_fcm"

type PayloadSendFcm struct {
	To      string    `json:"to"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Company int32     `json:"company"`
	Data    *DataNoti `json:"data"`
}

type DataNoti struct {
	Order   *int32 `json:"order"`
	Service *int32 `json:"service"`
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

	data, err := json.Marshal(payload.Data)
	if err != nil {
		log.Error().Err(err).Msg("failed to json marshal data")
	}
	payloadNoti := db.CreateNotificationParams{
		Type:    "SERVICE",
		Topic:   payload.To,
		Title:   payload.Title,
		Content: payload.Body,
		IsRead:  false,
		Company: sql.NullInt32{
			Int32: payload.Company,
			Valid: true,
		},
		Data: sql.NullString{
			String: string(data),
			Valid:  payload.Data != nil && err == nil,
		},
	}
	noti, err := processor.store.CreateNotification(ctx, payloadNoti)
	if err != nil {
		log.Error().Str("channel", payload.To).Msg("failed to create notification record db")
	}

	err = processor.client.SendMessage(payload.To, payload.Title, payload.Body)
	if err != nil {
		log.Error().Str("channel", payload.To).Msg("can't send fcm message")
		return err
	}

	ws := url.URL{Scheme: "ws", Host: ":8000", Path: fmt.Sprintf("/websocket/%d", payload.Company)}
	log.Info().Str("Connecting to %s", ws.String()).Msg("WS")

	conn, _, err := websocket.DefaultDialer.Dial(ws.String(), nil)
	if err != nil {
		log.Err(err).Msg("dial")
	}
	defer conn.Close()

	message, _ := json.Marshal(noti)
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Err(err).Msg("write")
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("channel", payload.To).
		Msg("processor task")
	return nil
}
