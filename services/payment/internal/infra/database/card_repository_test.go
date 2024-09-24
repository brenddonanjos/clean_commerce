package database

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCardRepositorySuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	cardEntity := &entity.Card{
		CardName:      "Card Name",
		Number:        "1234567890123456",
		HolderName:    "John Doe",
		ExpirityMonth: 12,
		ExpirityYear:  2024,
		CVV:           123,
		UserID:        1,
		AddressID:     1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(
		"INSERT INTO cards (card_name, number, holder_name, expirity_month, expirity_year, cvv, user_id, billing_address_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")).
		ExpectExec().
		WithArgs(
			cardEntity.CardName,
			cardEntity.Number,
			cardEntity.HolderName,
			cardEntity.ExpirityMonth,
			cardEntity.ExpirityYear,
			cardEntity.CVV,
			cardEntity.UserID,
			cardEntity.AddressID,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
		).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate the insertion with ID = 1

	cardRepository := NewCardRepository(sqlxDB)
	card, err := cardRepository.Save(cardEntity)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.NotNil(t, card.ID)
	assert.Equal(t, "Card Name", card.CardName)
	assert.Equal(t, "1234567890123456", card.Number)
	assert.Equal(t, "John Doe", card.HolderName)
	assert.Equal(t, int32(12), card.ExpirityMonth)
	assert.Equal(t, int32(2024), card.ExpirityYear)
	assert.Equal(t, int32(123), card.CVV)

}

func TestCardRepositoryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	cardEntity := &entity.Card{
		Number:        "1234567890123456",
		HolderName:    "John Doe",
		ExpirityMonth: 0,
		ExpirityYear:  2024,
		CVV:           0,
		UserID:        1,
		AddressID:     1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(
		"INSERT INTO cards (card_name, number, holder_name, expirity_month, expirity_year, cvv, user_id, billing_address_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")).
		ExpectExec().
		WithArgs(
			cardEntity.Number,
			cardEntity.HolderName,
			cardEntity.ExpirityMonth,
			cardEntity.ExpirityYear,
			cardEntity.CVV,
			cardEntity.UserID,
			cardEntity.AddressID,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
		).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate the insertion with ID = 1

	cardRepository := NewCardRepository(sqlxDB)
	card, err := cardRepository.Save(cardEntity)
	assert.NotNil(t, err)
	assert.Nil(t, card)
}
