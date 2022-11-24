package view_models

import (
	"mini-bank/src/models"
)

type TransactionListVm struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	CreditAccount int     `json:"credit_account"`
	DebitAccount  int     `json:"debit_account"`
	CreatedAt     string  `json:"created_at"`
}

type TransactionDetailVm struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	CreditAccount int     `json:"credit_account"`
	DebitAccount  int     `json:"debit_account"`
	CreatedAt     string  `json:"created_at"`
}

type TransactionVm struct {
	List   []TransactionListVm `json:"list_transaction"`
	Detail TransactionDetailVm `json:"detail_transaction"`
}

func NewTransactionVm() TransactionVm {
	return TransactionVm{}
}

func (vm TransactionVm) BuildList(transactions []models.Transaction) (res []*TransactionListVm) {
	if len(transactions) > 0 {
		for _, transaction := range transactions {
			res = append(res, &TransactionListVm{
				TransactionID: transaction.TransactionID,
				Amount:        transaction.Amount,
				CreditAccount: transaction.CreditAccount,
				DebitAccount:  transaction.DebitAccount,
				CreatedAt:     transaction.CreatedAt,
			})
		}
	}

	return res
}

func (vm TransactionVm) BuildDetail(transaction models.Transaction) (res *TransactionDetailVm) {
	nilTransaction := models.Transaction{}
	if transaction != nilTransaction {
		res = &TransactionDetailVm{
			TransactionID: transaction.TransactionID,
			Amount:        transaction.Amount,
			CreditAccount: transaction.CreditAccount,
			DebitAccount:  transaction.DebitAccount,
			CreatedAt:     transaction.CreatedAt,
		}
	}

	return res
}
