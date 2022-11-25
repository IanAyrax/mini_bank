package query

import "mini-bank/src/models"

type IAccountRepository interface {
	List() (res []models.Account, count int64, err error)
	Detail(id string) (res models.Account, err error)
}
