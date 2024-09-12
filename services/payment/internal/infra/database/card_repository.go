package database

import (
	"database/sql"
	"time"

	"github.com/brenddonanjos/payment_grpc_rabbitmq/services/payment/internal/entity"
)

type CardRepository struct {
	Db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (cr *CardRepository) Save(card *entity.Card) (*entity.Card, error) {
	stmt, err := cr.Db.Prepare("INSERT INTO cards (number, holder_name, expirity_month, expirity_year, cvv, user_id, address_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()

	result, err := stmt.Exec(
		card.Number,
		card.HolderName,
		card.ExpirityMonth,
		card.ExpirityYear,
		card.CVV,
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
	card.ID = int32(id)

	return card, nil
}
