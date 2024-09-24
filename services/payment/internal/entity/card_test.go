package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	card, err := NewCard("Card Name", "1234567890123456", "John Doe", 12, 2024, 123, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, "Card Name", card.CardName)
	assert.Equal(t, "1234567890123456", card.Number)
	assert.Equal(t, "John Doe", card.HolderName)
	assert.Equal(t, int32(12), card.ExpirityMonth)
	assert.Equal(t, int32(2024), card.ExpirityYear)
	assert.Equal(t, int32(123), card.CVV)
	assert.Equal(t, int32(1), card.UserID)
	assert.Equal(t, int32(1), card.AddressID)
}

func TestCardInvalid(t *testing.T) {
	card, err := NewCard("", "1234567890123456", "John Doe", 12, 2024, 123, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrCardNameIsRequired, err)

	card, err = NewCard("Card Name", "", "John Doe", 12, 2024, 123, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrNumberIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "", 12, 2024, 123, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrHolderNameIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "John Doe", 0, 2024, 123, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrExpirityMonthIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "John Doe", 12, 0, 123, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrExpirityYearIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "John Doe", 12, 2024, 0, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrCVVIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "John Doe", 12, 2024, 123, 0, 1)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrUserIDIsRequired, err)

	card, err = NewCard("Card Name", "1234567890123456", "John Doe", 12, 2024, 123, 1, 0)
	assert.NotNil(t, err)
	assert.Nil(t, card)
	assert.Equal(t, ErrAddressIDIsRequired, err)
}
