package services

import (
	"mini-bank/src/requests"
	"mini-bank/src/view_models"
)

type ITransactionService interface {
	Create(req *requests.TransactionRequest) (res *view_models.TransactionDetailVm, err error)
	List(req *requests.FilterTransactionRequest) (res []*view_models.TransactionListVm, err error)
	Detail(id string) (res *view_models.TransactionDetailVm, err error)
	Update(id string, req *requests.TransactionRequest) (err error)
	Delete()
}
