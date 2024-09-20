package entity

import (
	"context"
	"database/sql"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
)

type AddressRepositoryInterface interface {
	CreateAddress(ctx context.Context, arg database.CreateAddressParams) (sql.Result, error)
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (sql.Result, error)
}
