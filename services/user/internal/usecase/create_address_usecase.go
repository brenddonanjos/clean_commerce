package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/entity"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
	"github.com/brenddonanjos/clean_commerce/services/user/pkg/utils"
)

type AddressInputDTO struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	UserId  int32  `json:"user_id"`
}

type AddressOutputDTO struct {
	ID        int32     `json:"id"`
	Street    string    `json:"street"`
	Number    string    `json:"number"`
	ZipCode   string    `json:"zip_code"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	UserId    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAddressUseCase struct {
	AddressRepository entity.AddressRepositoryInterface
}

func NewCreateAddressUseCase(addressRepository entity.AddressRepositoryInterface) *CreateAddressUseCase {
	return &CreateAddressUseCase{AddressRepository: addressRepository}
}

func (uc *CreateAddressUseCase) Execute(input AddressInputDTO, ctx context.Context) (*AddressOutputDTO, error) {
	fmt.Println("Use case: Saving address...")
	cuurentDate := time.Now()
	address := &database.CreateAddressParams{
		Street:    utils.StringToSqlNullString(input.Street),
		Number:    utils.StringToSqlNullString(input.Number),
		ZipCode:   utils.StringToSqlNullString(input.ZipCode),
		City:      utils.StringToSqlNullString(input.City),
		Country:   utils.StringToSqlNullString(input.Country),
		State:     utils.StringToSqlNullString(input.State),
		UserID:    input.UserId,
		CreatedAt: utils.TimeToSqlNullTime(cuurentDate),
		UpdatedAt: utils.TimeToSqlNullTime(cuurentDate),
	}
	sqlResult, err := uc.AddressRepository.CreateAddress(ctx, *address)
	if err != nil {
		return nil, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &AddressOutputDTO{
		ID:      int32(id),
		Street:  address.Street.String,
		Number:  address.Number.String,
		ZipCode: address.ZipCode.String,
		City:    address.City.String,
		Country: address.Country.String,
		State:   address.State.String,
		UserId:  address.UserID,
	}, nil
}
