package db

import "context"

type CreateUserTxArgs struct {
	Username      string
	UserId        string
	PermissionIds []int64
	UserIds       []string
	CreatedByIds  []string
	AfterCreate   func(err error) error
}

type CreateUserTxResponse struct {
	Message string `json:"message"`
}

func (store *Store) CreateUserTx(ctx context.Context, args CreateUserTxArgs) (CreateUserTxResponse, error) {
	err := store.execTx(ctx, func(q *Queries) error {
		_, err := q.CreateUser(ctx, CreateUserParams{
			ID:       args.UserId,
			Username: args.Username,
		})

		if err != nil {
			return err
		}

		err = q.CreateUserPermissions(ctx, CreateUserPermissionsParams{
			UserID:       args.UserIds,
			PermissionID: args.PermissionIds,
			CreatedbyID:  args.CreatedByIds,
		})

		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		err = args.AfterCreate(err)

		return nil
	})

	return CreateUserTxResponse{
		Message: "User successfully created",
	}, err
}
