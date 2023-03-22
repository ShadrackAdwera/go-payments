package db

import "context"

type TxApproveRequest struct {
	ID     int64          `json:"id"`
	Status ApprovalStatus `json:"status"`
}

type TxApproveResponse struct {
	Request     Request     `json:"request"`
	UserPayment UserPayment `json:"user_payment"`
}

// Approve Request and Make new entry to join table
func (s *Store) ApprovePaymentRequest(ctx context.Context, txApproveRequest TxApproveRequest) (TxApproveResponse, error) {
	var txApproveResponse TxApproveResponse
	err := s.execTx(ctx, func(q *Queries) error {
		req, err := q.UpdateRequest(ctx, UpdateRequestParams{
			ID: txApproveRequest.ID,
			Status: NullApprovalStatus{
				ApprovalStatus: txApproveRequest.Status,
				Valid:          true,
			},
		})

		if err != nil {
			return err
		}

		txApproveResponse.UserPayment, err = q.CreateUserPayment(ctx, CreateUserPaymentParams{
			RequestID: req.ID,
			ClientID:  req.PaidToID,
		})

		if err != nil {
			return err
		}

		txApproveResponse.Request = req
		return nil
	})
	if err != nil {
		return TxApproveResponse{}, err
	}
	return txApproveResponse, nil
}
