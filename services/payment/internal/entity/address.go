package entity

import "time"

type Address struct {
	ID        int32     `json:"id" db:"id"`
	Street    string    `json:"street" db:"street"`
	Number    string    `json:"number" db:"number"`
	ZipCode   string    `json:"zip_code" db:"zip_code"`
	City      string    `json:"city" db:"city"`
	Country   string    `json:"country" db:"country"`
	State     string    `json:"state" db:"state"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
