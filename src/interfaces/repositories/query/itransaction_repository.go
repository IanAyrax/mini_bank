package query

import "mini-bank/src/models"

type ITransactionRepository interface {
	List(creditAccount, debitAccount string) (res []models.Transaction, count int64, err error)
	Detail(id string) (res models.Transaction, err error)
}
