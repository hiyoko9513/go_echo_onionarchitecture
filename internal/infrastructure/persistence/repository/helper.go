package repository

import (
	"fmt"

	"hiyoko-echo/internal/pkg/mypubliclib/ent"
)

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
