package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/ShadrackAdwera/go-payments/utils"
)

// fetch daraja token
// check if expired
// create daraja token
// delete expired token
// save to db
// return new token

type DarajaTxResponse struct {
	DarajaToken DarajaToken `json:"daraja_token"`
}

func (store *Store) DarajaTokenTx(ctx context.Context) (DarajaTxResponse, error) {
	var tknResponse DarajaTxResponse

	err := store.execTx(ctx, func(q *Queries) error {
		tkn, err := q.GetDarajaToken(ctx)

		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
		}

		if err == sql.ErrNoRows {
			darajaRes, err := utils.RequestDarajaToken()
			if err != nil {
				return err
			}
			// request for token / store in the db return
			newTkn, err := q.CreateDarajaToken(ctx, CreateDarajaTokenParams{
				AccessToken: darajaRes.AccessToken,
				ExpiresAt:   darajaRes.ExpiresAt,
			})

			if err != nil {
				return err
			}
			tknResponse.DarajaToken = newTkn
			return nil
		}

		if time.Now().After(tkn.ExpiresAt) {
			// request for token / store in the db then delete the current access token
			darajaRes, err := utils.RequestDarajaToken()
			if err != nil {
				return err
			}

			// request for token / store in the db return
			newTkn, err := q.CreateDarajaToken(ctx, CreateDarajaTokenParams{
				AccessToken: darajaRes.AccessToken,
				ExpiresAt:   darajaRes.ExpiresAt,
			})

			if err != nil {
				return err
			}
			tknResponse.DarajaToken = newTkn

			err = q.DeleteDarajaToken(ctx, tkn.ID)
			if err != nil {
				return err
			}
			return nil
		}

		tknResponse.DarajaToken = tkn
		return nil
	})
	return tknResponse, err
}
