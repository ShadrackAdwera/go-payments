package db

import (
	"context"
	"database/sql"
)

type TxStore interface {
	Querier
	ApproveRequestTx(ctx context.Context, args ApproveRequestTxRequest) (ApproveRequestTxResponse, error)
	DarajaTokenTx(ctx context.Context) (DarajaTxResponse, error)
	CreateUserTx(ctx context.Context, args CreateUserTxArgs) (CreateUserTxResponse, error)
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) TxStore {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(context context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(context, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}
