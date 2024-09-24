package service

import (
	"context"

	"github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/grpc/pb_payment"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BillingAddressService struct {
	pb_payment.UnimplementedBillingAddressServiceServer
	CreateBillingAddressUseCase usecase.CreateBillingAddressUseCase
}

func NewBillingAddressService(createBillingAddressUseCase *usecase.CreateBillingAddressUseCase) *BillingAddressService {
	return &BillingAddressService{
		CreateBillingAddressUseCase: *createBillingAddressUseCase,
	}
}

func (b *BillingAddressService) CreateBillingAddress(
	ctx context.Context,
	req *pb_payment.CreateBillingAddressRequest) (*pb_payment.BillingAddressResponse, error) {
	createBillingAddressDTO := usecase.BillingAddressInputDTO{
		Street:  req.Street,
		Number:  req.Number,
		ZipCode: req.ZipCode,
		City:    req.City,
		State:   req.State,
		Country: req.Country,
		UserID:  req.UserId,
	}

	outputDTO, err := b.CreateBillingAddressUseCase.Execute(createBillingAddressDTO)
	if err != nil {
		return nil, err
	}

	return &pb_payment.BillingAddressResponse{
		Id:        outputDTO.ID,
		Street:    outputDTO.Street,
		Number:    outputDTO.Number,
		ZipCode:   outputDTO.ZipCode,
		City:      outputDTO.City,
		Country:   outputDTO.Country,
		State:     outputDTO.State,
		UserId:    outputDTO.UserID,
		CreatedAt: timestamppb.New(outputDTO.CreatedAt),
		UpdatedAt: timestamppb.New(outputDTO.UpdatedAt),
	}, nil
}
