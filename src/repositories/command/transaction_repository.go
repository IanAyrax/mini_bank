package command

import (
	"fmt"
	"mini-bank/src/interfaces/repositories/command"
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewCommandTransactionRepository(db *gorm.DB) command.ITransactionRepository {
	return &TransactionRepository{DB: db}
}

func (TransactionRepository) Create(transaction models.Transaction, tx *gorm.DB) (res models.Transaction, err error) {
	err = tx.Create(&transaction).Error
	if err != nil {
		return res, err
	}

	fmt.Println(transaction)

	return transaction, nil
}

func (TransactionRepository) Update(transaction models.Transaction, tx *gorm.DB) (res models.Transaction, err error) {
	err = tx.Updates(&transaction).Error
	if err != nil {
		return res, err
	}

	return transaction, nil
}

func (TransactionRepository) Delete(transaction models.Transaction, tx *gorm.DB) (err error) {
	err = tx.Delete(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
