package models

type Account struct {
	AccountNo int    `gorm:"primaryKey;"`
	Name      string `gorm:"unique;type:varchar(50)"`
	Balance   float64
	CreatedAt string `gorm:"type:varchar(20)"`
}
