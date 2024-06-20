package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Account Services
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connectionString := "user=postgres dbname=go-minibank sslmode=disable"
	// connectionString := "postgres://go-minibank:password@localhost/go-minibank?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (store *PostgresStore) Init() error {
	return store.createAccountTable()
}

func (store *PostgresStore) createAccountTable() error {
	query := `create table if not exists account(
			id serial primary key,
			first_name varchar(100),
			last_name varchar(100),
			number serial,
			encrypted_password varchar(100),
			balance serial,
			created_at timestamp
	)`
	_, err := store.db.Exec(query)
	return err
}

// CreateAccount implements Storage.
func (store *PostgresStore) CreateAccount(account *Account) error {
	query := `insert into account (first_name, last_name, number, balance, created_at) values ($1, $2, $3, $4, $5)`
	response, err := store.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", response)

	return nil
}

func (store *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := store.db.Query("select * from account")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// GetAccountByID implements Storage.
func (store *PostgresStore) GetAccountByID(accountID int) (*Account, error) {
	rows, err := store.db.Query("select * from account where id = $1", accountID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", accountID)
}

// DeleteAccount implements Storage.
func (store *PostgresStore) DeleteAccount(int) error {
	panic("unimplemented")
}

// UpdateAccount implements Storage.
func (store *PostgresStore) UpdateAccount(*Account) error {
	panic("unimplemented")
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
