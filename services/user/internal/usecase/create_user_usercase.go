package usecase

import (
	"context"
	"log"
	"time"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/entity"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
	"github.com/brenddonanjos/clean_commerce/services/user/pkg/utils"
)

type UserInputDTO struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}

type UserOutputDTO struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository entity.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

func (uc *CreateUserUseCase) Execute(input UserInputDTO, ctx context.Context) (*UserOutputDTO, error) {
	log.Println("Use case: Saving user...")
	cuurentDate := time.Now()
	user := &database.CreateUserParams{
		Name:      input.Name,
		Email:     input.Email,
		Username:  input.Username,
		Password:  input.Password,
		BirthDate: utils.TimeToSqlNullTime(input.BirthDate),
		CreatedAt: utils.TimeToSqlNullTime(cuurentDate),
		UpdatedAt: utils.TimeToSqlNullTime(cuurentDate),
	}

	sqlResult, err := uc.UserRepository.CreateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	response := &UserOutputDTO{
		ID:        int32(id),
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		BirthDate: user.BirthDate.Time,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
	return response, nil
}
