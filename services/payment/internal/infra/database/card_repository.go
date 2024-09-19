package database

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	Db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (cr *CardRepository) Save(card *entity.Card) (*entity.Card, error) {
	fmt.Println("Saving card...")
	stmt, err := cr.Db.Prepare("INSERT INTO cards (card_name, number, holder_name, expirity_month, expirity_year, cvv, user_id, address_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()

	result, err := stmt.Exec(
		card.CardName,
		card.Number,
		card.HolderName,
		card.ExpirityMonth,
		card.ExpirityYear,
		card.CVV,
		card.UserID,
		card.AddressID,
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
