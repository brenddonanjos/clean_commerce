package entity

import "time"

type Card struct {
	ID            int32     `json:"id" db:"id"`
	CardName      string    `json:"card_name" db:"card_name"`
	Number        string    `json:"number" db:"number"`
	HolderName    string    `json:"holder_name" db:"holder_name"`
	ExpirityMonth int32     `json:"expirity_month" db:"expirity_month"`
	ExpirityYear  int32     `json:"expirity_year" db:"expirity_year"`
	CVV           int32     `json:"cvv" db:"cvv"`
	UserID        int32     `json:"user_id" db:"user_id"`
	AddressID     int32     `json:"address_id" db:"address_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`
}
