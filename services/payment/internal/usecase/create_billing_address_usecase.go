package usecase

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
)

type BillingAddressInputDTO struct {
	Street  string
	Number  string
	ZipCode string
	City    string
	Country string
	State   string
	UserID  int32
}

type BillingAddressOutputDTO struct {
	ID        int32
	Street    string
	Number    string
	ZipCode   string
	City      string
	Country   string
	State     string
	UserID    int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateBillingAddressUseCase struct {
	BillingAddressRepository entity.BillingAddressRepositoryInterface
}

func NewCreateBillingAddressUseCase(
	billingAddressRepository entity.BillingAddressRepositoryInterface,
) *CreateBillingAddressUseCase {
	return &CreateBillingAddressUseCase{
		BillingAddressRepository: billingAddressRepository,
	}
}

func (c *CreateBillingAddressUseCase) Execute(input BillingAddressInputDTO) (*BillingAddressOutputDTO, error) {
	fmt.Println("Use case: Saving billing address...")

	address, err := entity.NewBillingAddress(
		input.Street,
		input.Number,
		input.ZipCode,
		input.City,
		input.Country,
		input.State,
		input.UserID,
	)
	if err != nil {
		return nil, err
	}

	address, err = c.BillingAddressRepository.Save(address)
	if err != nil {
		return nil, err
	}

	return &BillingAddressOutputDTO{
		ID:        address.ID,
		Street:    address.Street,
		Number:    address.Number,
		ZipCode:   address.ZipCode,
		City:      address.City,
		Country:   address.Country,
		State:     address.State,
		UserID:    address.UserID,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}, nil
}
