package requests

type TransactionRequest struct {
	Amount        float64 `form:"amount" json:"amount"`
	CreditAccount int     `form:"credit_account" json:"credit_account" validate:"required"`
	DebitAccount  int     `form:"debit_account" json:"debit_account" validate:"required"`
}

type FilterTransactionRequest struct {
	CreditAccount int `form:"credit_account" json:"credit_account"`
	DebitAccount  int `form:"debit_account" json:"debit_account"`
}
