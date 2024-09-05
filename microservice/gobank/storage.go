package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(string) error
	UpdateAccount(*Account) error
	GetAccountByID(string) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostegresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostegresStore, error) {

	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {

		return nil, err
	}

	return &PostegresStore{
		db: db,
	}, nil

}

func (p *PostegresStore) Init() error {

	if err := p.CreateAccountTable(); err != nil {

		return err
	}

	return nil
}

func (p *PostegresStore) CreateAccountTable() error {

	query := `CREATE TABLE IF NOT EXISTS ACCOUNT (

	ID SERIAL PRIMARY KEY,
	FIRST_NAME VARCHAR(50),
	LAST_NAME VARCHAR(50),
	NUMBER SERIAL,
	BALANCE SERIAL,
	CREATED_AT TIMESTAMP
	
	)`

	_, err := p.db.Query(query)

	return err

}

func (p *PostegresStore) CreateAccount(acc *Account) error {

	query := `INSERT INTO ACCOUNT
	 (FIRST_NAME,LAST_NAME,NUMBER,BALANCE,CREATED_AT) 
	 VALUES ($1,$2,$3,$4,$5)`

	resp, err := p.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	fmt.Println(resp)

	if err != nil {

		return err
	}

	return nil
}

func (p *PostegresStore) DeleteAccount(id string) error {

	query := "DELETE FROM ACCOUNT WHERE ID = $1"

	_, err := p.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostegresStore) UpdateAccount(*Account) error {
	return nil
}

func (p *PostegresStore) GetAccountByID(id string) (*Account, error) {

	query := "SELECT * FROM ACCOUNT WHERE ID = $1"

	rows, err := p.db.Query(query, id)
	if err != nil {

		return nil, err
	}

	account := new(Account)
	for rows.Next() {
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

	}

	return account, nil
}

func (p *PostegresStore) GetAccounts() ([]*Account, error) {

	rows, err := p.db.Query("SELECT * FROM ACCOUNT")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {

			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}
