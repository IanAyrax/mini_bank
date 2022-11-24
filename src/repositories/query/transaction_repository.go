package query

import (
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewQueryTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (repo TransactionRepository) List(creditAccount, debitAccount string) (res []models.Transaction, count int64, err error) {
	db := repo.DB

	if creditAccount != "" {
		db = db.Where("credit_account = ?", creditAccount)
	}

	if debitAccount != "" {
		db = db.Where("debit_account = ?", debitAccount)
	}

	err = db.Find(&res).Error
	if err != nil {
		return res, count, err
	}

	err = db.Find(&models.Transaction{}).Count(&count).Error
	if err != nil {
		return res, count, err
	}

	return res, count, nil
}

func (repo TransactionRepository) Detail(id string) (res models.Transaction, err error) {
	db := repo.DB
	err = db.Where("transaction_id = ?", id).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
