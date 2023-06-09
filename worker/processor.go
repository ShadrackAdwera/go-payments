package worker

import (
	"context"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskProcessor interface {
	TaskProcessPayment(ctx context.Context, task *asynq.Task) error
}

type PaymentTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewTaskServer(opts asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(opts, asynq.Config{
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Err(err).Str("task_type", task.Type()).Bytes("payload", task.Payload()).Msg("error processing task . . ")
		}),
	})
	return &PaymentTaskProcessor{
		server, store,
	}
}
