package repository

import (
	"context"
	"hiyoko-echo/domain/repository"
	"hiyoko-echo/ent"
	"hiyoko-echo/ent/util"
	"hiyoko-echo/infrastructure/database"
)

type userRepository struct {
	Conn *database.EntClient
}

func NewUserRepository(conn *database.EntClient) repository.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) List(ctx context.Context) ([]*ent.User, error) {
	u, err := r.Conn.User.Query().Limit(10).Offset(0).All(ctx)
	return u, err
}

func (r *userRepository) Get(ctx context.Context, id util.ID) (*ent.User, error) {
	u, err := r.Conn.User.Get(ctx, id)
	return u, err
}

func (r *userRepository) Create(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := r.Conn.User.Create().
		SetName(u.Name).
		Save(ctx)
	return u, err
}

func (r *userRepository) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := u.Update().SetName(u.Name).Save(ctx)
	return u, err
}

func (r *userRepository) Delete(ctx context.Context, id util.ID) error {
	u, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	err = r.Conn.User.DeleteOne(u).Exec(ctx)
	return err
}
