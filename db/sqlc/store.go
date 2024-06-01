package simplebank

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all function to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// create a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// exectx executes a function wihtin a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err : %v, rb err : %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// transfertxparam contains the input parameter of transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// transfertxresult is the result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entriy"`
	ToEntry     Entry    `json:"to_entry"`
}
