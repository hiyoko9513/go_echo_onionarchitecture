package repository

import (
	"context"

	"hiyoko-echo/domain/repository"
	"hiyoko-echo/ent/migrate"
	"hiyoko-echo/infrastructure/database"
)

type tableRepository struct {
	Conn *database.EntClient
}

func NewTableRepository(conn *database.EntClient) repository.TableRepository {
	return &tableRepository{conn}
}

func (r *tableRepository) Migrate(ctx context.Context) error {
	err := r.Conn.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}

// todo 利用する関数にctxを利用する系にする
func (r *tableRepository) TruncateAll(ctx context.Context) error {
	sqlclient := r.Conn.DB()
	_, err := sqlclient.Exec("SET FOREIGN_KEY_CHECKS=0;")
	if err != nil {
		return err
	}

	var truncateQuery string
	test, err := sqlclient.Query("SELECT CONCAT('TRUNCATE TABLE ', GROUP_CONCAT(CONCAT('`',table_name,'`')),';') AS statement FROM information_schema.tables WHERE table_schema = 'hiyoko' AND table_name LIKE '%';")
	if err != nil {
		return err
	}
	test.Next()
	err = test.Scan(&truncateQuery)
	if err != nil {
		return err
	}
	err = test.Close()
	if err != nil {
		return err
	}
	_, err = sqlclient.Exec(truncateQuery)
	if err != nil {
		return err
	}

	_, err = sqlclient.Exec("SET FOREIGN_KEY_CHECKS=1;")
	if err != nil {
		return err
	}
	return nil
}
