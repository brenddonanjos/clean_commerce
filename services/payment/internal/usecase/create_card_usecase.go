package usecase

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/entity"
)

type CardInputDTO struct {
	CardName    string `json:"card_name"`
	Number      string `json:"number"`
	HolderName  string `json:"holder_name"`
	ExpiryMonth int32  `json:"expiry_month"`
	ExpiryYear  int32  `json:"expiry_year"`
	CVV         int32  `json:"cvv"`
	UserId      int32  `json:"user_id"`
	AddressId   int32  `json:"address_id"`
}

type CardOutputDTO struct {
	ID          int32     `json:"id"`
	CardName    string    `json:"card_name"`
	Number      string    `json:"number"`
	HolderName  string    `json:"holder_name"`
	ExpiryMonth int32     `json:"expiry_month"`
	ExpiryYear  int32     `json:"expiry_year"`
	CVV         int32     `json:"cvv"`
	UserId      int32     `json:"user_id"`
	AddressId   int32     `json:"address_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCardUseCase struct {
	CardRepository entity.CardRepositoryInterface
}

func NewCreateCardUseCase(cardRepository entity.CardRepositoryInterface) *CreateCardUseCase {
	return &CreateCardUseCase{CardRepository: cardRepository}
}

func (uc *CreateCardUseCase) Execute(input CardInputDTO) (*CardOutputDTO, error) {
	fmt.Println("Use case: Saving card...")
	card, err := entity.NewCard(
		input.CardName,
		input.Number,
		input.HolderName,
		input.ExpiryMonth,
		input.ExpiryYear,
		input.CVV,
		input.UserId,
		input.AddressId,
	)
	if err != nil {
		return nil, err
	}

	card, err = uc.CardRepository.Save(card)
	if err != nil {
		return nil, err
	}

	return &CardOutputDTO{
		ID:          card.ID,
		CardName:    card.CardName,
		Number:      card.Number,
		HolderName:  card.HolderName,
		ExpiryMonth: card.ExpirityMonth,
		ExpiryYear:  card.ExpirityYear,
		CVV:         card.CVV,
		UserId:      card.UserID,
		AddressId:   card.AddressID,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}, nil
}
