package service

import (
	"context"

	pb "github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/pb_user"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AddressService struct {
	pb.UnimplementedAddressServiceServer
	CreateAddressUseCase usecase.CreateAddressUseCase
}

func NewAddressService(createAddressUseCase *usecase.CreateAddressUseCase) *AddressService {
	return &AddressService{CreateAddressUseCase: *createAddressUseCase}
}

func (s *AddressService) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.AddressResponse, error) {
	createAddressDTO := usecase.AddressInputDTO{
		Street:  req.Street,
		Number:  req.Number,
		ZipCode: req.ZipCode,
		City:    req.City,
		Country: req.Country,
		State:   req.State,
		UserId:  req.UserId,
	}

	address, err := s.CreateAddressUseCase.Execute(createAddressDTO, ctx)
	if err != nil {
		return nil, err
	}

	return &pb.AddressResponse{
		Id:        address.ID,
		Street:    address.Street,
		Number:    address.Number,
		ZipCode:   address.ZipCode,
		City:      address.City,
		Country:   address.Country,
		State:     address.State,
		UserId:    address.UserId,
		CreatedAt: timestamppb.New(address.CreatedAt),
		UpdatedAt: timestamppb.New(address.UpdatedAt),
	}, nil
}
