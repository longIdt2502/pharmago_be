package woker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const TaskUploadImageVariant = "task:upload_image_variant"

type PayloadUploadImageVariant struct {
	Image []byte `json:"image"`
	Id    int32  `json:"id"`
}

func (distributor *RedisTaskDistributor) DistributorUploadImageVariant(ctx context.Context, payload *PayloadUploadImageVariant, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskUploadImageVariant, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task:, %w", err)
	}
	log.Info().
		Str("type", task.Type()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueue task:upload_image_variant")
	return nil
}

func (processor *RedisTaskProcessor) ProcessorUploadImageVariant(ctx context.Context, task *asynq.Task) error {
	var payload PayloadUploadImageVariant
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}
	url, err := processor.b2Bucket.UploadFileToB2(payload.Image)
	if err != nil {
		return fmt.Errorf("failed to upload image to b2: %e", err)
	}
	log.Info().Str("url:", url).Msg("processor task:upload_image_variant")

	media, err := processor.store.CreateMedia(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to record media: %e", err)
	}

	_, err = processor.store.CreateVariantMedia(ctx, db.CreateVariantMediaParams{
		Variant: payload.Id,
		Media:   media.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to record db: %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Str("id variant", string(payload.Id)).
		Msg("processor task:upload_image_variant")
	return nil
}

const TaskUploadImageProduct = "task:upload_image_product"

type PayloadUploadImageProduct struct {
	Image []byte `json:"image"`
	Id    int32  `json:"id"`
}

func (distributor *RedisTaskDistributor) DistributorUploadImageProduct(ctx context.Context, payload *PayloadUploadImageProduct, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskUploadImageProduct, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task:, %w", err)
	}
	log.Info().
		Str("type", task.Type()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueue task:upload_image_product")
	return nil
}

func (processor *RedisTaskProcessor) ProcessorUploadImageProduct(ctx context.Context, task *asynq.Task) error {
	var payload PayloadUploadImageProduct
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	url, err := processor.b2Bucket.UploadFileToB2(payload.Image)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("failed to upload image to b2: %e", err))
	}

	media, err := processor.store.CreateMedia(ctx, url)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("failed to record media: %e", err))
	}

	_, err = processor.store.CreateProductMedia(ctx, db.CreateProductMediaParams{
		Product: payload.Id,
		Media:   media.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create product media: %w", err)
	}
	log.Info().
		Str("type", task.Type()).
		Str("id product", string(payload.Id)).
		Msg("processor task:upload_image_product")
	return nil
}
