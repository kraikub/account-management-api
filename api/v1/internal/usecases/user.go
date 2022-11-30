package usecases

import (
	"context"
	"fmt"

	"github.com/kraikub/account-management-api/api/v1/internal/dtos"
	"github.com/kraikub/account-management-api/api/v1/internal/repositories"
)

type UserUseCase interface {
	FindUserWithUserId(ctx context.Context, uid string) (dtos.UserDTO, error)
	UpdateUserWithUserId(ctx context.Context, uid string, uDTO dtos.UserDTO) error
}

type userUseCase struct {
	userRepository repositories.UserRepository
}

func CreateUserUseCase(userRepository repositories.UserRepository) userUseCase {
	return userUseCase{
		userRepository: userRepository,
	}
}

func (u userUseCase) FindUserWithUserId(ctx context.Context, uid string) (dtos.UserDTO, error) {
	userE, err := u.userRepository.FindUserWithUserId(ctx, uid)
	if err != nil {
		return dtos.UserDTO{}, err
	}
	return dtos.UserDTO{}.FromEntity(userE), nil
}

func (u userUseCase) UpdateUserWithUserId(ctx context.Context, uid string, uDTO dtos.UserDTO) error {
	fmt.Println(uid, uDTO)
	return u.userRepository.UpdateUserWithUserId(ctx, uid, uDTO.ToUpdateableEntity())
}
