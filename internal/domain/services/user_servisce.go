package services

import (
	"context"

	"hiyoko-echo/internal/pkg/ent"
	"hiyoko-echo/internal/pkg/ent/util"
)

type UserRepository interface {
	List(ctx context.Context) ([]*ent.User, error)
	Get(ctx context.Context, id util.ID) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	Update(ctx context.Context, user *ent.User) (*ent.User, error)
	Delete(ctx context.Context, user *ent.User) error
}
