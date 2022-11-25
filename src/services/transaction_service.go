package services

import (
	"fmt"
	"mini-bank/src/interfaces/services"
	"mini-bank/src/models"
	"mini-bank/src/repositories/command"
	"mini-bank/src/repositories/query"
	"mini-bank/src/requests"
	"mini-bank/src/view_models"
	"time"
)

type TransactionService struct {
	*Contract
}

func NewTransactionService(contract *Contract) services.ITransactionService {
	return &TransactionService{Contract: contract}
}

func (transactionService TransactionService) Create(req *requests.TransactionRequest) (res *view_models.TransactionDetailVm, err error) {
	db := transactionService.DB
	repo := command.NewCommandTransactionRepository(db)

	model := models.Transaction{
		Amount:        req.Amount,
		CreditAccount: req.CreditAccount,
		DebitAccount:  req.DebitAccount,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, err := repo.Create(model, db)

	res = view_models.NewTransactionVm().BuildDetail(transaction)

	return res, nil
}

func (transactionService TransactionService) List(req *requests.FilterTransactionRequest) (res []*view_models.TransactionListVm, err error) {
	db := transactionService.DB
	repo := query.NewQueryTransactionRepository(db)

	creditAccount := ""
	debitAccount := ""

	if req != nil {
		if req.CreditAccount != 0 {
			creditAccount = fmt.Sprint(req.CreditAccount)
		}

		if req.DebitAccount != 0 {
			debitAccount = fmt.Sprint(req.DebitAccount)
		}
	}

	transactions, _, err := repo.List(creditAccount, debitAccount)
	if err != nil {
		return res, err
	}

	res = view_models.NewTransactionVm().BuildList(transactions)

	return res, nil
}

func (transactionService TransactionService) Detail(id string) (res *view_models.TransactionDetailVm, err error) {
	db := transactionService.DB
	repo := query.NewQueryTransactionRepository(db)
	transaction, err := repo.Detail(id)
	if err != nil {
		return res, err
	}

	res = view_models.NewTransactionVm().BuildDetail(transaction)

	return res, nil
}

func (transactionService TransactionService) Update(id string, req *requests.TransactionRequest) (err error) {
	db := transactionService.DB
	repo := command.NewCommandTransactionRepository(db)

	model := models.Transaction{
		TransactionID: id,
		Amount:        req.Amount,
		CreditAccount: req.CreditAccount,
		DebitAccount:  req.DebitAccount,
	}

	_, err = repo.Update(model, db)
	if err != nil {
		return err
	}

	return nil
}

func (transactionService TransactionService) Delete() {

}
