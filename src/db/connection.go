package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

// DB is struct for database instance
type DB struct {
	Conn *gorm.DB
}

// NewDB is constructor to create db instance
func NewDB() (*DB, error) {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	ssl := "disable"
	env := os.Getenv("ENV")
	if env == "production" {
		ssl = "require"
	}

	authDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		ssl,
		os.Getenv("DB_PASS"),
	)

	conn, err := gorm.Open("postgres", authDB)
	if err != nil {
		return nil, err
	}

	logging.Success("Success to connect database")
	return &DB{Conn: conn}, nil
}

// Get is method to get connection
func (d DB) Get() *gorm.DB {
	return d.Conn
}
