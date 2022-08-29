package repository

import (
	"github.com/chenningg/hermitboard-api/db"
	"github.com/oklog/ulid/v2"
)

type AccountRepository interface {
	GetOne(id ulid.ULID) db.Account
	GetAll() []db.Account
	New() db.Account
}

type accountRepository struct {
}
