package main

import (
	"database/sql"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(uuid.UUID) error
	UpdateAccount(*Account) error
	GetAccountByID(uuid.UUID) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=schlucht sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil

}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {

	query := `CREATE TABLE IF NOT EXISTS Account(
		id uuid PRIMARY KEY ,
		first_name varchar(255),
		last_name varchar(255),
		number serial,
		balance serial,
		create_at TIMESTAMP)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) GetAccounts() ([]*Account, error) {
	query := "SELECT * FROM account"
	var accounts = []*Account{}

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, nil
	}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgressStore) CreateAccount(account *Account) error {
	query := `
		INSERT INTO account 
		(id, first_name, last_name, number, balance, create_at)
		VALUES
		($1,$2,$3,$4,$5, $6)
	`
	_, err := s.db.Query(query,
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) DeleteAccount(id uuid.UUID) error {

	return nil
}

func (s *PostgressStore) UpdateAccount(account *Account) error {

	return nil
}

func (s *PostgressStore) GetAccountByID(id uuid.UUID) (*Account, error) {
	var account *Account
	var err error
	query := `
		SELECT * FROM account WHERE ID = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		account, err = scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
	}

	return account, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	var account Account
	if err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &account, nil
}
