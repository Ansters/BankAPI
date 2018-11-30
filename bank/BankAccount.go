package bank

type BankAccount struct {
	id            int     `json:"id"`
	userID        int     `json:"user_id"`
	accountNumber int     `json:"account_number"`
	name          string  `json:"name"`
	balance       float64 `json:"balance"`
}
