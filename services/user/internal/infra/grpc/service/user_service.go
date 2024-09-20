package service

import (
	"context"
	"fmt"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/pb_user"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	pb_user.UnimplementedUserServiceServer
	CreateUserUseCase usecase.CreateUserUseCase
}

func NewUserService(createUserUseCase *usecase.CreateUserUseCase) *UserService {
	return &UserService{CreateUserUseCase: *createUserUseCase}
}

func (us *UserService) CreateUser(ctx context.Context, req *pb_user.CreateUserRequest) (*pb_user.UserResponse, error) {
	fmt.Println("Use case: Saving user...")
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}
	createUserDTO := usecase.UserInputDTO{
		Name:      req.Name,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		BirthDate: birthDate,
	}

	user, err := us.CreateUserUseCase.Execute(createUserDTO, ctx)
	if err != nil {
		return nil, err
	}

	return &pb_user.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		BirthDate: timestamppb.New(user.BirthDate),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil

}
