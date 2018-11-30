package userapi

import "database/sql"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BankAccount struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	AccountNumber int     `json:"account_number"`
	Balance       float64 `json:"balance"`
}

type Service struct {
	DB *sql.DB
}

func (s *Service) FindByID(id int) (*User, error) {
	stmt := `SELECT id, first_name, last_name FROM Users WHERE id=$1;`
	row := s.DB.QueryRow(stmt, id)
	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) All() ([]User, error) {
	stmt := `SELECT id, first_name, last_name FROM Users;`
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Service) CreateUser(user *User) error {
	stmt := `INSERT INTO Users (first_name, last_name) VALUES($1, $2) RETURNING id`
	row := s.DB.QueryRow(stmt, user.FirstName, user.LastName)
	err := row.Scan(&user.ID)
	return err
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
