package entity

import (
	"context"
	"database/sql"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (sql.Result, error)
}
