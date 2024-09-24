package database

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBillingAddressRepositorySuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// Set a fixed time for testing
	fixedTime := time.Now().Truncate(time.Second)

	addressEntity := &entity.BillingAddress{
		Street:    "Street",
		Number:    "123",
		ZipCode:   "12345",
		City:      "City",
		Country:   "Country",
		State:     "State",
		UserID:    1,
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(
		"INSERT INTO billing_addresses (street, number, zip_code, city, country, state, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")).
		ExpectExec().
		WithArgs(
			addressEntity.Street,
			addressEntity.Number,
			addressEntity.ZipCode,
			addressEntity.City,
			addressEntity.Country,
			addressEntity.State,
			addressEntity.UserID,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
		).WillReturnResult(sqlmock.NewResult(1, 1))

	billingAddressRepository := NewBillingAddressRepository(sqlxDB)
	address, err := billingAddressRepository.Save(addressEntity)

	assert.NoError(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, addressEntity.Street, address.Street)
	assert.Equal(t, addressEntity.Number, address.Number)
	assert.Equal(t, addressEntity.ZipCode, address.ZipCode)
	assert.Equal(t, addressEntity.City, address.City)
	assert.Equal(t, addressEntity.Country, address.Country)
	assert.Equal(t, addressEntity.State, address.State)
	assert.Equal(t, addressEntity.UserID, address.UserID)
}

func TestBillingAddressRepositoryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	addressEntity := &entity.BillingAddress{
		Number:    "123",
		ZipCode:   "12345",
		City:      "City",
		Country:   "Country",
		State:     "State",
		UserID:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mock.ExpectPrepare(regexp.QuoteMeta(
		"INSERT INTO billing_addresses (street, number, zip_code, city, country, state, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")).
		ExpectExec().
		WithArgs(
			addressEntity.Number,
			addressEntity.ZipCode,
			addressEntity.City,
			addressEntity.Country,
			addressEntity.State,
			addressEntity.UserID,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
		).WillReturnResult(sqlmock.NewResult(1, 1))

	billingAddressRepository := NewBillingAddressRepository(sqlxDB)
	address, err := billingAddressRepository.Save(addressEntity)

	assert.Error(t, err)
	assert.Nil(t, address)
	assert.NotNil(t, err)
}
