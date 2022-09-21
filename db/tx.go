package db

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
)

// WithTx wraps a function and executes it using a database transaction.
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	// Create a transactional client.
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("WithTx(): %v", err)
	}

	// If there is a panicking goroutine then intercept it to rollback the transaction.
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Otherwise, call the closured function to run the transaction.
	if err := fn(tx); err != nil {
		// If the closured functions returns an error, rollback the transaction.
		if rerr := tx.Rollback(); rerr != nil {
			// There was an error rolling back the transaction. Show both the function error and the rollback error.
			err = fmt.Errorf("WithTx(): %v: %v", rerr, err)
		}
		return err
	}

	// Otherwise the function is good to go and we can commit the transaction.
	if err := tx.Commit(); err != nil {
		// There was an error committing the transaction.
		return fmt.Errorf("WithTx(): %v", err)
	}
	return nil
}
