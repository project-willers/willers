package config

import (
	"fmt"
	"os"
)

func Port() string {
	return os.Getenv("PORT")
}

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	) + "?parseTime=true&loc=Asia%2FTokyo&collation=utf8mb4_bin"
}

func Secret() string {
	return os.Getenv("SECRET")
}
