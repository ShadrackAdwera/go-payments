package db

import (
	"context"
	"database/sql"
	"time"
)

type ApproveRequestTxRequest struct {
	ID             int64          `json:"id"`
	ApprovalStatus ApprovalStatus `json:"approval_status"`
	AfterApproval  func(clientId int64, amount int64, userPaymentId int64, status ApprovalStatus) error
}

type ApproveRequestTxResponse struct {
	Request     Request     `json:"request"`
	UserPayment UserPayment `json:"user_payment"`
}

func (s *Store) ApproveRequestTx(ctx context.Context, args ApproveRequestTxRequest) (ApproveRequestTxResponse, error) {
	var response ApproveRequestTxResponse

	err := s.execTx(ctx, func(q *Queries) error {
		req, err := q.UpdateRequest(ctx, UpdateRequestParams{
			ID: args.ID,
			Status: NullApprovalStatus{
				ApprovalStatus: args.ApprovalStatus,
				Valid:          true,
			},
			ApprovedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})

		if err != nil {
			return err
		}

		p, err := q.CreateUserPayment(ctx, CreateUserPaymentParams{
			RequestID: sql.NullInt64{
				Int64: req.ID,
				Valid: true,
			},
			ClientID: sql.NullInt64{
				Int64: req.PaidToID,
				Valid: true,
			},
			Status: PaidStatusNotPaid,
		})

		if err != nil {
			return err
		}

		err = args.AfterApproval(req.PaidToID, req.Amount, p.ID, args.ApprovalStatus)

		if err != nil {
			return err
		}

		response.Request = req
		response.UserPayment = p

		return nil
	})

	return response, err
}
