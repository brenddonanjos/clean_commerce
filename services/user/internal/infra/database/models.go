// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type CommerceUserUser struct {
	ID        int32
	Name      string
	Username  string
	Email     string
	Password  string
	BirthDate sql.NullTime
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
