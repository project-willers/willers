package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Database *sql.DB
)

func Initdb() {
	var err error
	Database, err = sql.Open("mysql", "root:example@tcp(willers-mysql:3306)/willers?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("Database Connect error: ", err)
	}
	Database.SetConnMaxLifetime(time.Minute * 3)
	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(10)

	for i := 0; i < 4; i++ {
		if err = Database.PingContext(context.Background()); err == nil {
			log.Println("Connected to Database:", err)
			break
		}
		log.Println("Ping error:", err)
		time.Sleep(time.Second * 10)
	}
}
