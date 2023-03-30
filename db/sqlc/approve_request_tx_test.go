package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApproveRequestTx(t *testing.T) {
	n := 5

	approveReqTxChan := make(chan ApproveRequestTxResponse)
	errChan := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			req := CreateRandomRequest(t)

			res, err := testQuery.ApproveRequestTx(context.Background(), ApproveRequestTxRequest{
				ID:             req.ID,
				ApprovalStatus: req.Status,
			})

			approveReqTxChan <- res
			errChan <- err
		}()
	}

	for j := 0; j < n; j++ {
		resp := <-approveReqTxChan
		err := <-errChan

		require.NoError(t, err)
		require.NotEmpty(t, resp)
	}
}
