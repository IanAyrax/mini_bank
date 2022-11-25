package command

import (
	"mini-bank/src/interfaces/repositories/command"
	"mini-bank/src/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewCommandAccountRepository(db *gorm.DB) command.IAccountRepository {
	return &AccountRepository{DB: db}
}

func (AccountRepository) Create(account models.Account, tx *gorm.DB) (res models.Account, err error) {
	err = tx.Create(&account).Error
	if err != nil {
		return res, err
	}
	return account, nil
}

func (AccountRepository) Update(account models.Account, tx *gorm.DB) (res models.Account, err error) {
	err = tx.Updates(&account).Error
	if err != nil {
		return res, err
	}

	return account, nil
}

func (AccountRepository) Delete(account models.Account, tx *gorm.DB) (err error) {
	err = tx.Delete(&account).Error
	if err != nil {
		return err
	}

	return nil
}
