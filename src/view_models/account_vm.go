package view_models

import (
	"mini-bank/src/models"
)

type AccountVm struct {
	List   []*AccountListVm `json:"list_account"`
	Detail AccountListVm    `json:"detail_account"`
}

type AccountListVm struct {
	AccountNo int     `json:"account_no"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}

type AccountListNoVm struct {
	AccountNo int `json:"account_no"`
}

type AccountDetailVm struct {
	AccountNo int     `json:"account_no"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}

func NewAccountVm() AccountVm {
	return AccountVm{}
}

func (vm AccountVm) BuildList(accounts []models.Account) (res []*AccountListVm) {
	if len(accounts) > 0 {
		for _, account := range accounts {
			res = append(res, &AccountListVm{
				AccountNo: account.AccountNo,
				Name:      account.Name,
				Balance:   account.Balance,
				CreatedAt: account.CreatedAt,
			})
		}
	}

	return res
}

func (vm AccountVm) BuildListNoFromVmList(accounts []*AccountListVm) (res []*AccountListNoVm) {
	if len(accounts) > 0 {
		for _, account := range accounts {
			res = append(res, &AccountListNoVm{
				AccountNo: account.AccountNo,
			})
		}
	}

	return res
}

func (vm AccountVm) BuildDetail(account models.Account) (res *AccountDetailVm) {
	nilAccount := models.Account{}
	if account != nilAccount {
		res = &AccountDetailVm{
			AccountNo: account.AccountNo,
			Name:      account.Name,
			Balance:   account.Balance,
			CreatedAt: account.CreatedAt,
		}
	}

	return res
}
