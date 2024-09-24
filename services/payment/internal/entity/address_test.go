package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressSuccess(t *testing.T) {

	address, err := NewBillingAddress("Street", "123", "ZipCode", "City", "Country", "State", 1)
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, "Street", address.Street)
	assert.Equal(t, "123", address.Number)
	assert.Equal(t, "ZipCode", address.ZipCode)
	assert.Equal(t, "City", address.City)
	assert.Equal(t, "Country", address.Country)
	assert.Equal(t, "State", address.State)
	assert.Equal(t, int32(1), address.UserID)
}

func TestAddressFail(t *testing.T) {
	address, err := NewBillingAddress("", "123", "ZipCode", "City", "Country", "State", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrStreetIsRequired, err)

	address, err = NewBillingAddress("Street", "", "ZipCode", "City", "Country", "State", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrNumberIsRequired, err)

	address, err = NewBillingAddress("Street", "123", "", "City", "Country", "State", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrZipCodeIsRequired, err)

	address, err = NewBillingAddress("Street", "123", "ZipCode", "", "Country", "State", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrCityIsRequired, err)

	address, err = NewBillingAddress("Street", "123", "ZipCode", "City", "", "State", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrCountryIsRequired, err)

	address, err = NewBillingAddress("Street", "123", "ZipCode", "City", "Country", "", 1)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrStateIsRequired, err)

	address, err = NewBillingAddress("Street", "123", "ZipCode", "City", "Country", "State", 0)
	assert.NotNil(t, err)
	assert.Nil(t, address)
	assert.Equal(t, ErrUserIDIsRequired, err)
}
