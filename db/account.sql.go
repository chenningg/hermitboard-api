// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: account.sql

package db

import (
	"context"
	"database/sql"

	hbtype "github.com/chenningg/hermitboard-api/hbtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
    id, auth_id, nickname, email, password
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, auth_id, nickname, email, password
`

type CreateAccountParams struct {
	ID       hbtype.ULID
	AuthID   sql.NullString
	Nickname string
	Email    string
	Password sql.NullString
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount,
		arg.ID,
		arg.AuthID,
		arg.Nickname,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Nickname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :one
DELETE FROM account
WHERE id = $1
RETURNING id, auth_id, nickname, email, password
`

func (q *Queries) DeleteAccount(ctx context.Context, id hbtype.ULID) (Account, error) {
	row := q.db.QueryRow(ctx, deleteAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Nickname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, auth_id, nickname, email, password FROM account
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id hbtype.ULID) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Nickname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, auth_id, nickname, email, password FROM account
ORDER BY email
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.AuthID,
			&i.Nickname,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account
SET auth_id = $2,
nickname = $3,
email = $4,
password = $5
WHERE id = $1
RETURNING id, auth_id, nickname, email, password
`

type UpdateAccountParams struct {
	ID       hbtype.ULID
	AuthID   sql.NullString
	Nickname string
	Email    string
	Password sql.NullString
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount,
		arg.ID,
		arg.AuthID,
		arg.Nickname,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Nickname,
		&i.Email,
		&i.Password,
	)
	return i, err
}