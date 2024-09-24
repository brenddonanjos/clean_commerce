package entity

import (
	"errors"
	"time"
)

var (
	ErrCardNameIsRequired      = errors.New("card name is required")
	ErrNumberIsRequired        = errors.New("number is required")
	ErrHolderNameIsRequired    = errors.New("holder name is required")
	ErrExpirityMonthIsRequired = errors.New("expirity month is required")
	ErrExpirityYearIsRequired  = errors.New("expirity year is required")
	ErrCVVIsRequired           = errors.New("cvv is required")
	ErrUserIDIsRequired        = errors.New("user id is required")
	ErrAddressIDIsRequired     = errors.New("address id is required")
)

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

func NewCard(
	cardName string,
	number string,
	holderName string,
	expirityMonth int32,
	expirityYear int32,
	cvv int32,
	userID int32,
	addressID int32) (*Card, error) {

	card := &Card{
		CardName:      cardName,
		Number:        number,
		HolderName:    holderName,
		ExpirityMonth: expirityMonth,
		ExpirityYear:  expirityYear,
		CVV:           cvv,
		UserID:        userID,
		AddressID:     addressID,
	}

	err := card.validate()
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (card *Card) validate() error {
	if card.CardName == "" {
		return ErrCardNameIsRequired
	}
	if card.Number == "" {
		return ErrNumberIsRequired
	}
	if card.HolderName == "" {
		return ErrHolderNameIsRequired
	}
	if card.ExpirityMonth == 0 {
		return ErrExpirityMonthIsRequired
	}
	if card.ExpirityYear == 0 {
		return ErrExpirityYearIsRequired
	}
	if card.CVV == 0 {
		return ErrCVVIsRequired
	}
	if card.UserID == 0 {
		return ErrUserIDIsRequired
	}
	if card.AddressID == 0 {
		return ErrAddressIDIsRequired
	}

	return nil
}
