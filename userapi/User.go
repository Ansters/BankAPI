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
	AccountNumber string  `json:"account_number"`
	Name          string  `json:"name"`
	Balance       float64 `json:"balance"`
}

type Amount struct {
	Amount float64 `JSON:"amount"`
	From   int     `JSON:"from"`
	To     int     `JSON:"to"`
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
	stmt := `SELECT id, first_name, last_name FROM Users ORDER BY id ASC;`
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
	stmt := `UPDATE Users SET first_name=$1, last_name=$2 WHERE id=$3;`
	_, err := s.DB.Exec(stmt, user.FirstName, user.LastName, user.ID)
	return err
}

func (s *Service) Delete(id int) error {
	stmt := `DELETE FROM Users WHERE id=$1;`
	_, err := s.DB.Exec(stmt, id)
	return err
}

func (s *Service) CreateBank(b *BankAccount) error {
	stmt := `INSERT INTO BankAccount(user_id, account_number, name, balance) VALUES($1, $2, $3, 0.0) RETURNING id`
	row := s.DB.QueryRow(stmt, b.ID, b.AccountNumber, b.Name)
	err := row.Scan(&b.ID)
	return err
}

func (s *Service) AllBankByID(id int) ([]BankAccount, error) {
	stmt := `SELECT id, user_id, account_number, name, balance FROM BankAccount WHERE user_id=$1;`
	rows, err := s.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	var bankAccs []BankAccount
	for rows.Next() {
		var b BankAccount
		err := rows.Scan(&b.ID, &b.UserID, &b.AccountNumber, &b.Name, &b.Balance)
		if err != nil {
			return nil, err
		}
		bankAccs = append(bankAccs, b)
	}
	return bankAccs, nil
}

func (s *Service) DeleteBankByID(id int) error {
	stmt := `DELETE FROM BankAccount WHERE id=$1;`
	_, err := s.DB.Exec(stmt, id)
	return err
}

func (s *Service) Withdraw(id int, amount Amount) error {
	stmt := `UPDATE BankAccount SET balance = balance-$1 WHERE id=$2;`
	_, err := s.DB.Exec(stmt, amount.Amount, id)
	return err
}

func (s *Service) Deposit(id int, amount Amount) error {
	stmt := `UPDATE BankAccount SET balance = balance+$1 WHERE id=$2;`
	_, err := s.DB.Exec(stmt, amount.Amount, id)
	return err
}

func (s *Service) Transfer(fromID int, toID int, amount float64) error {
	return nil
}
