package repository

import (
	"context"
)

type TableRepository interface {
	Migrate(ctx context.Context) error
	TruncateAll(ctx context.Context) error
	//Seeder(ctx context.Context) error
}
