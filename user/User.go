package user

import "database/sql"

type User struct {
	id        int    `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}

type BankAccount struct {
	id            int     `json:"id"`
	userID        int     `json:"user_id"`
	accountNumber int     `json:"account_number"`
	name          string  `json:"name"`
	balance       float64 `json:"balance"`
}

type Service struct {
	DB *sql.DB
}

func (s *Service) FindByID(id int) (*User, error) {
	return nil, nil
}

func (s *Service) All() ([]User, error) {
	return nil, nil
}

func (s *Service) CreateUser(user *User) error {
	return nil
}

func (s *Service) Update(user *User) error {
	return nil
}

func (s *Service) Delete(user *User) error {
	return nil
}

func (s *Service) CreateBank(id int, b BankAccount) error {
	return nil
}

func (s *Service) AllBankByID(id int) ([]BankAccount, error) {
	return nil, nil
}

func (s *Service) DeleteBankByID(id int) error {
	return nil
}

func (s *Service) Withdraw(id int, amount float64) error {
	return nil
}

func (s *Service) Deposit(id int, amount float64) error {
	return nil
}

func (s *Service) Transfer(fromID int, toID int, amount float64) error {
	return nil
}
