package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributePayment(ctx context.Context, payload *PaymentPayload, opts ...asynq.Option) error
}

type PaymentTaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(clientOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(clientOpts)
	return &PaymentTaskDistributor{
		client,
	}
}
