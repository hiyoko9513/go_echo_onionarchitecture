package usecase

import (
	"context"

	"hiyoko-echo/domain/repository"
	"hiyoko-echo/internal/pkg/mypubliclib/ent"
	"hiyoko-echo/internal/pkg/mypubliclib/ent/util"
)

type UserUseCase interface {
	GetUsers(ctx context.Context) ([]*ent.User, error)
	GetUser(ctx context.Context, id util.ID) (*ent.User, error)
	CreateUser(ctx context.Context, user *ent.User) (*ent.User, error)
	UpdateUser(ctx context.Context, id util.ID) (*ent.User, error)
	DeleteUser(ctx context.Context, id util.ID) error
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return u.UserRepository.List(ctx)
}

func (u *userUseCase) GetUser(ctx context.Context, id util.ID) (*ent.User, error) {
	return u.UserRepository.Get(ctx, id)
}

func (u *userUseCase) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return u.UserRepository.Create(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, id util.ID) (*ent.User, error) {
	user, err := u.UserRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.UserRepository.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id util.ID) error {
	user, err := u.UserRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(ctx, user)
}
