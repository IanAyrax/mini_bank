package requests

type AccountRequest struct {
	Name    string  `form:"name" json:"name" validate:"required"`
	Balance float64 `form:"balance" json:"balance"`
}

type FilterQueryAccountRequest struct {
	Detail bool `form:"detail" json:"detail"`
}
