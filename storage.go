package main

import "database/sql"

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

// CreateAccount implements Storage.
func (store *PostgresStore) CreateAccount(*Account) error {
	panic("unimplemented")
}

// DeleteAccount implements Storage.
func (store *PostgresStore) DeleteAccount(int) error {
	panic("unimplemented")
}

// GetAccountByID implements Storage.
func (store *PostgresStore) GetAccountByID(int) (*Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements Storage.
func (store *PostgresStore) UpdateAccount(*Account) error {
	panic("unimplemented")
}

func NewPostgresStore() (*PostgresStore, error) {

	connectionString := "user=postgres dbname=gominibank sslmode=disable"

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
