package command

import (
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(transaction models.Transaction, tx *gorm.DB) (res models.Transaction, err error)
	Update(transaction models.Transaction, tx *gorm.DB) (res models.Transaction, err error)
	Delete(transaction models.Transaction, tx *gorm.DB) (err error)
}
