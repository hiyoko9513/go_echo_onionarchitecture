package repository

import (
	"context"
	"fmt"

	"hiyoko-echo/pkg/mypubliclib/ent"
)

func (r *tableRepository) Seed(ctx context.Context) error {
	tx, err := r.conn.Tx(ctx)
	if err != nil {
		err = fmt.Errorf("failed to seed; error: %v", err)
		return err
	}

	var usersInputs []ent.User
	for i := 0; i < 100; i++ {
		usersInputs = append(usersInputs, ent.User{
			Name: fmt.Sprintf("user+%d", i),
		})
	}

	bulk := make([]*ent.UserCreate, len(usersInputs))
	for i, input := range usersInputs {
		bulk[i] = tx.User.Create().
			SetName(input.Name)
	}
	_, err = tx.User.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		err = fmt.Errorf("failed to seed; error: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = rollback(tx, err)
		err = fmt.Errorf("failed to seed; error: %v", err)
		return err
	}
	return nil
}
