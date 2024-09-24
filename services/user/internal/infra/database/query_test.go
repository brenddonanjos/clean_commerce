package database

import (
	"context"
	"testing"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/user/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestUserRepository(t *testing.T) {
	db, err := setupTestDB()
	assert.Nil(t, err)
	defer db.Close()

	birthDate := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	userEntity := &CreateUserParams{
		Name:      "Name",
		Username:  "Username",
		Email:     "Email",
		Password:  "Password",
		BirthDate: utils.TimeToSqlNullTime(birthDate),
	}

	userRepository := New(db)
	sqlRes, err := userRepository.CreateUser(context.Background(), *userEntity)
	assert.Nil(t, err)
	assert.NotNil(t, sqlRes)
	id, err := sqlRes.LastInsertId()
	assert.Nil(t, err)
	assert.NotNil(t, id)

}

func setupTestDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`ATTACH DATABASE ':memory:' AS commerce_user`)
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS commerce_user.users (
		id INTEGER PRIMARY KEY, 
		name TEXT NOT NULL, 
		username TEXT NOT NULL, 
		email TEXT NOT NULL, 
		password TEXT NOT NULL, 
		birth_date DATE, 
		created_at DATETIME, 
		updated_at DATETIME, 
		deleted_at DATETIME
	);`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
