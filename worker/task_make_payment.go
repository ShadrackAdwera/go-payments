package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type PaymentPayload struct {
	ClientID      int64 `json:"client_id"`
	Amount        int64 `json:"amount"`
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

func (processor *PaymentTaskProcessor) TaskProcessPayment(ctx context.Context, task *asynq.Task) error {

	var payload PaymentPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshall json %w", asynq.SkipRetry)
	}

	// get client details using client id
	clientData, err := processor.store.GetClient(ctx, payload.ClientID)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("client was not found %w", asynq.SkipRetry)
		}
		return fmt.Errorf("an error occured - 1 %w", err)
	}

	// send payment

	tkn, err := processor.store.DarajaTokenTx(ctx)

	if err != nil {
		return fmt.Errorf("an error occured daraja token - 2 %w", err)
	}

	_, err = utils.MakeMobileMoneyPayment(clientData.Phone, payload.Amount, "Organization Name", tkn.DarajaToken.AccessToken)

	if err != nil {
		return fmt.Errorf("an error occured while makein the payment - 2 %w", err)
	}

	// update user payment to paid
	up, err := processor.store.UpdateUserPayment(ctx, db.UpdateUserPaymentParams{
		ID:     payload.UserPaymentID,
		Status: db.PaidStatusPaid,
	})

	if err != nil {
		return fmt.Errorf("an error occured - 3 %w", err)
	}

	log.Info().Str("client_name", clientData.Name).Str("payment made to", clientData.Email).
		Str("preferred_payment", string(clientData.PreferredPaymentType)).
		Int64("user_payment_id", up.ID).
		Msg("Client info found")

	return nil
}
