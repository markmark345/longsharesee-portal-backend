// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO "users" (
    email, "password", phone, region, gender, title, first_name, last_name, age
) VALUES (
    'user01@gmail.com', 'password', '1234567890', 'thai', 'M', 'Mr.', 'testf', 'testl', 24
)
`

func (q *Queries) CreateUser(ctx context.Context) error {
	_, err := q.db.Exec(ctx, createUser)
	return err
}
