// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, dollar_1 pgtype.UUID) error
	GetUsers(ctx context.Context, arg GetUsersParams) (User, error)
}

var _ Querier = (*Queries)(nil)