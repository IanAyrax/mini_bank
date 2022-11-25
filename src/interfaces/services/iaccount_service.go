package services

import (
	"mini-bank/src/requests"
	"mini-bank/src/view_models"
)

type IAccountService interface {
	Create(req *requests.AccountRequest) (res *view_models.AccountDetailVm, err error)
	List(req *requests.FilterQueryAccountRequest) (res []*view_models.AccountListVm, err error)
	Detail(id string) (res *view_models.AccountDetailVm, err error)
	Update(id string, req *requests.AccountRequest) (err error)
	Delete()
}
