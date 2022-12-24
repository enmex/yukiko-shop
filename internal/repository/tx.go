package repository

import "yukiko-shop/internal/repository/ent"

func rollback(tx *ent.Tx, cause error) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return cause
}
