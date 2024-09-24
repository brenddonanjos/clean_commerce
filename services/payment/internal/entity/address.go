package entity

import (
	"errors"
	"time"
)

var (
	ErrStreetIsRequired           = errors.New("street is required")
	ErrAddressNumberIsRequired    = errors.New("card number is required")
	ErrZipCodeIsRequired          = errors.New("zip code is required")
	ErrCityIsRequired             = errors.New("city is required")
	ErrCountryIsRequired          = errors.New("country is required")
	ErrStateIsRequired            = errors.New("state is required")
	ErrBillingAddressIDIsRequired = errors.New("billing address id is required")
)

type BillingAddress struct {
	ID        int32     `json:"id" db:"id"`
	Street    string    `json:"street" db:"street"`
	Number    string    `json:"number" db:"number"`
	ZipCode   string    `json:"zip_code" db:"zip_code"`
	City      string    `json:"city" db:"city"`
	Country   string    `json:"country" db:"country"`
	State     string    `json:"state" db:"state"`
	UserID    int32     `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewBillingAddress(Street, Number, ZipCode, City, Country, State string, UserID int32) (*BillingAddress, error) {
	address := &BillingAddress{
		Street:  Street,
		Number:  Number,
		ZipCode: ZipCode,
		City:    City,
		Country: Country,
		State:   State,
		UserID:  UserID,
	}

	err := address.validate()
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (a *BillingAddress) validate() error {
	if a.Street == "" {
		return ErrStreetIsRequired
	}
	if a.Number == "" {
		return ErrNumberIsRequired
	}
	if a.ZipCode == "" {
		return ErrZipCodeIsRequired
	}
	if a.City == "" {
		return ErrCityIsRequired
	}
	if a.Country == "" {
		return ErrCountryIsRequired
	}
	if a.State == "" {
		return ErrStateIsRequired
	}
	if a.UserID == 0 {
		return ErrUserIDIsRequired
	}

	return nil
}
