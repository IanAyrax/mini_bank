package models

type Transaction struct {
	TransactionID string `gorm:"primaryKey;type:varchar(36);default:(uuid())"`
	Amount        float64
	CreditAccount int
	DebitAccount  int
	CreatedAt     string   `gorm:"type:varchar(20)"`
	Credit        *Account `gorm:"foreignKey:CreditAccount;constraint:OnUpdate:Cascade,OnDelete:Cascade;"`
	Debit         *Account `gorm:"foreignKey:DebitAccount;constraint:OnUpdate:Cascade,OnDelete:Cascade;"`
}
