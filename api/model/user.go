package model

import (
	"context"
	"database/sql"
	"log"
	"net/mail"
	"time"

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
	Name     string `json:"name"`
	Password string `json:"password" validate:"required,gte=8,lt=50"`
}

func FindUser(u *LoginInfo) (*Account, error) {
	var result *sql.Row
	if validEmail(u.Name) {
		result = db.Database.QueryRowContext(context.Background(), "SELECT * FROM accounts WHERE email=?", u.Name)
	} else {
		result = db.Database.QueryRowContext(context.Background(), "SELECT * FROM accounts WHERE name=?", u.Name)
	}

	user := &Account{}
	if err := result.Scan(user.Name, user.Email, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(u *Account) (*Account, error) {
	if _, err := FindUser(&LoginInfo{Name: u.Name}); err != nil {
		log.Println(err)
		return u, err
	}
	result := db.Database.QueryRowContext(context.Background(), "INSERT INTO accounts(name, email, password) VALUE(?, ?, ?)", u.Name, u.Email, u.Password)
	if err := result.Scan(u.Name, u.Email, u.Password); err != nil {
		return u, err
	}
	return u, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
