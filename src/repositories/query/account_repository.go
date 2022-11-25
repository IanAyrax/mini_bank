package query

import (
	"mini-bank/src/interfaces/repositories/query"
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewQueryAccountRepository(db *gorm.DB) query.IAccountRepository {
	return &AccountRepository{DB: db}
}

func (repo AccountRepository) List() (res []models.Account, count int64, err error) {
	db := repo.DB

	err = db.Find(&res).Error
	if err != nil {
		return res, count, err
	}

	err = db.Find(&models.Account{}).Count(&count).Error
	if err != nil {
		return res, count, err
	}

	return res, count, nil
}

func (repo AccountRepository) Detail(id string) (res models.Account, err error) {
	db := repo.DB
	err = db.Where("account_no = ?", id).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
