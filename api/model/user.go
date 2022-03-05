package model

import (
	"context"
	"database/sql"
	"log"
	"net/mail"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"willers-api/db"
)

type UserInfo struct {
	ID       string `validate:"required,gte=3,lt=10"`
	Name     string `validate:"required,gte=1,lt=20"`
	Email    string `validate:"required,email"`
	password string `validate:"required,gte=8,lt=50"`

	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
	description string    `validate:"lt=200"`
}

type Account struct {
	Name     string `json:"name" validate:"required,gte=2,lt=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lt=50"`
}

type LoginInfo struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,gte=8,lt=50"`
}

func FindUser(u *LoginInfo) (*Account, error) {
	var result *sql.Row
	if validEmail(u.Name) {
		result = db.Database.QueryRowContext(context.Background(), "SELECT * FROM accounts WHERE email=?", u.Name)
	} else {
		result = db.Database.QueryRowContext(context.Background(), "SELECT * FROM accounts WHERE name=?", u.Name)
	}

	account := &Account{}
	if err := result.Scan(&account.Name, &account.Email, &account.Password); err != nil {
		return nil, err
	}
	return account, nil
}

func CreateUser(u *Account) (*Account, error) {
	if _, err := FindUser(&LoginInfo{Name: u.Name}); err == nil {
		return u, err
	}
	log.Println(u)
	insert, err := db.Database.Prepare("INSERT INTO accounts(name, email, password) VALUE(?, ?, ?)")
	if err != nil {
		return u, err
	}
	defer insert.Close()
	result, err := insert.ExecContext(context.Background(), u.Name, u.Email, u.Password)
	if err != nil {
		return u, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return u, err
	}
	log.Println(rowCnt)
	return u, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
