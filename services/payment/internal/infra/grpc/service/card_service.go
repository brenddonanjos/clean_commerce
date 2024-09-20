package service

import (
	"context"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/grpc/pb_payment"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CardService struct {
	pb_payment.UnimplementedCardServiceServer
	CreateCardUseCase usecase.CreateCardUseCase
}

func NewCardService(createCardUsecase *usecase.CreateCardUseCase) *CardService {
	return &CardService{
		CreateCardUseCase: *createCardUsecase,
	}
}

func (cs *CardService) CreateCard(ctx context.Context, request *pb_payment.CreateCardRequest) (*pb_payment.CardResponse, error) {
	createCardDTO := usecase.CardInputDTO{
		CardName:    request.CardName,
		Number:      request.Number,
		HolderName:  request.HolderName,
		ExpiryMonth: request.ExpiryMonth,
		ExpiryYear:  request.ExpiryYear,
		CVV:         request.Cvv,
		UserId:      request.UserId,
		AddressId:   request.AddressId,
	}

	outputDTO, err := cs.CreateCardUseCase.Execute(createCardDTO)
	if err != nil {
		return nil, err
	}

	response := &pb_payment.CardResponse{
		Id:          outputDTO.ID,
		CardName:    outputDTO.CardName,
		Number:      outputDTO.Number,
		HolderName:  outputDTO.HolderName,
		ExpiryMonth: outputDTO.ExpiryMonth,
		ExpiryYear:  outputDTO.ExpiryYear,
		Cvv:         outputDTO.CVV,
		UserId:      outputDTO.UserId,
		AddressId:   outputDTO.AddressId,
		CreatedAt:   timestamppb.New(outputDTO.CreatedAt),
		UpdatedAt:   timestamppb.New(outputDTO.UpdatedAt),
	}

	return response, nil
}
