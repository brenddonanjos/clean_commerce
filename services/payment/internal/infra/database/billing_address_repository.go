package database

import (
	"time"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BillingAddressRepository struct {
	db *sqlx.DB
}

func NewBillingAddressRepository(db *sqlx.DB) *BillingAddressRepository {
	return &BillingAddressRepository{db: db}
}

func (b *BillingAddressRepository) Save(address *entity.BillingAddress) (*entity.BillingAddress, error) {
	stmt, err := b.db.Prepare("INSERT INTO billing_addresses (street, number, zip_code, city, country, state, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	result, err := stmt.Exec(
		address.Street,
		address.Number,
		address.ZipCode,
		address.City,
		address.Country,
		address.State,
		address.UserID,
		currentTime,
		currentTime,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	address.ID = int32(id)
	return address, nil
}
