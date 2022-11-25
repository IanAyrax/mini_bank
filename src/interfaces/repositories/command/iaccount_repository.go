package command

import (
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	Create(account models.Account, tx *gorm.DB) (res models.Account, err error)
	Update(account models.Account, tx *gorm.DB) (res models.Account, err error)
	Delete(account models.Account, tx *gorm.DB) (err error)
}
