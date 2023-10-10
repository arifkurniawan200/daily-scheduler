package repository

import (
	"database/sql"
	"template/internal/model"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionHandler{db}
}

func (t TransactionHandler) CreateTransactionTx(tx *sql.Tx, transaction model.Transaction) error {
	query := createTransactionWithPromo
	if transaction.CampaignID == 0 {
		query = createTransactionWithoutPromo
		_, err := tx.Exec(query, transaction.UserID, transaction.ProductID, transaction.Total, transaction.Amount, transaction.Status)
		if err != nil {
			return err
		}
		return err
	}
	_, err := tx.Exec(query, transaction.UserID, transaction.ProductID, transaction.CampaignID, transaction.Total, transaction.Amount, transaction.Status)
	if err != nil {
		return err
	}
	return err
}
