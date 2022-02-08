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
	Name     string    `json:"name"`
	Password string    `json:"password" validate:"required,gte=8,lt=50"`
	CreateAt time.Time `json:"created_at"`
}

func FindUser(u *LoginInfo) (Account, error) {
	var (
		user   Account
		result *sql.Rows
	)
	if _, err := mail.ParseAddress(u.Name); err != nil {
		result, err = db.Database.QueryContext(context.Background(), "SELECT * FROM accounts WHERE name=?", u.Name)
	} else {
		result, err = db.Database.QueryContext(context.Background(), "SELECT * FROM accounts WHERE email=?", u.Name)
	}
	defer result.Close()
	for result.Next() {
		if err := result.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			return user, err
		}
		break
	}
	return user, nil
}

func CreateUser(u *Account) (*Account, error) {
	var result *sql.Rows
	if _, err := FindUser(&LoginInfo{Name: u.Name}); err != nil {
		log.Println(err)
		return u, err
	}
	result, err := db.Database.QueryContext(context.Background(), "INSERT INTO accounts(name, email, password) VALUE(?, ?, ?)", u.Name, u.Email, u.Password)
	if err != nil {
		return u, err
	}
	defer result.Close()
	for result.Next() {
		if err := result.Scan(&u.Name, &u.Email, &u.Password); err != nil {
			return u, err
		}
		break
	}
	return u, nil
}
