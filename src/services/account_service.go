package services

import (
	"mini-bank/src/interfaces/services"
	"mini-bank/src/models"
	"mini-bank/src/repositories/command"
	"mini-bank/src/repositories/query"
	"mini-bank/src/requests"
	"mini-bank/src/view_models"
	"strconv"
	"time"
)

type AccountService struct {
	*Contract
}

func NewAccountService(contract *Contract) services.IAccountService {
	return &AccountService{Contract: contract}
}

func (accountService AccountService) Create(req *requests.AccountRequest) (res *view_models.AccountDetailVm, err error) {
	db := accountService.DB
	repo := command.NewCommandAccountRepository(db)

	model := models.Account{
		Name:      req.Name,
		Balance:   req.Balance,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	account, err := repo.Create(model, db)
	if err != nil {
		return nil, err
	}

	res = view_models.NewAccountVm().BuildDetail(account)

	return res, nil
}

func (accountService AccountService) List(req *requests.FilterQueryAccountRequest) (res []*view_models.AccountListVm, err error) {
	db := accountService.DB
	repo := query.NewQueryAccountRepository(db)

	accounts, _, err := repo.List()
	if err != nil {
		return res, err
	}

	res = view_models.NewAccountVm().BuildList(accounts)

	return res, nil
}

func (accountService AccountService) Detail(id string) (res *view_models.AccountDetailVm, err error) {
	db := accountService.DB
	repo := query.NewQueryAccountRepository(db)
	account, err := repo.Detail(id)
	if err != nil {
		return res, err
	}

	res = view_models.NewAccountVm().BuildDetail(account)

	return res, nil
}

func (accountService AccountService) Update(id string, req *requests.AccountRequest) (err error) {
	db := accountService.DB
	repo := command.NewCommandAccountRepository(db)
	accountNo, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	model := models.Account{
		AccountNo: accountNo,
		Name:      req.Name,
		Balance:   req.Balance,
	}

	_, err = repo.Update(model, db)
	if err != nil {
		return err
	}

	return nil
}

func (accountService AccountService) Delete() {

}
