package usecase

import (
	"context"

	"github.com/ttakezawa/go-service-example/domain"
)

type UserUsecase struct {
	UserRepository domain.UserRepository `inject:""`
}

func (userUsecase *UserUsecase) Get(c context.Context, name string) (*domain.User, error) {
	return userUsecase.UserRepository.FindByName(name)
}
