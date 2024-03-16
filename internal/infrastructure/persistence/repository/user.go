package repository

import (
	"context"

	"hiyoko-echo/internal/domain/services"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/internal/pkg/mypubliclib/ent"
	"hiyoko-echo/internal/pkg/mypubliclib/ent/util"
)

type userRepository struct {
	conn *database.EntClient
}

func NewUserRepository(conn *database.EntClient) services.UserRepository {
	return &userRepository{conn}
}

// todo curdをentの共通関数とするような作りにする
// todo txの導入

func (r *userRepository) List(ctx context.Context) ([]*ent.User, error) {
	u, err := r.conn.User.Query().Limit(10).Offset(0).All(ctx)
	return u, err
}

func (r *userRepository) Get(ctx context.Context, id util.ID) (*ent.User, error) {
	u, err := r.conn.User.Get(ctx, id)
	return u, err
}

func (r *userRepository) Create(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := r.conn.User.Create().
		SetName(u.Name).
		Save(ctx)
	return u, err
}

func (r *userRepository) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := u.Update().SetName(u.Name).Save(ctx)
	return u, err
}

func (r *userRepository) Delete(ctx context.Context, u *ent.User) error {
	err := r.conn.User.DeleteOne(u).Exec(ctx)
	return err
}
