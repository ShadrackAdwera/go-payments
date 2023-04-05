package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type PaymentPayload struct {
	ClientID      int64 `json:"client_id"`
	Amount        int64 `json:"amount"`
	RequestID     int64 `json:"request_id"`
	UserPaymentID int64 `json:"user_payment_id"`
}

const TaskMakePayment = "task:make_payment"

func (distro *PaymentTaskDistributor) DistributePayment(ctx context.Context, payload *PaymentPayload, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("failed to marshall json body")
	}

	taskPay := asynq.NewTask(TaskMakePayment, jsonPayload, opts...)

	info, err := distro.client.EnqueueContext(ctx, taskPay)

	if err != nil {
		return fmt.Errorf("unable to enqueue task context : %w", err)
	}

	log.Info().
		Str("task_type", info.Type).
		Str("task_id", info.ID).
		Str("queue", info.Queue).
		Bytes("payload", jsonPayload).
		Int("max_retries", info.MaxRetry).
		Msg("task enqueued")

	return nil
}
