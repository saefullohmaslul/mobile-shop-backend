package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

// Conn is connection session
var Conn *gorm.DB

// DB is struct for database instance
type DB struct {
}

// NewDB is constructor to create db instance
func NewDB() (bool, error) {
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
		return false, err
	}
	conn.DB().SetMaxIdleConns(100)

	logging.Success("Success to connect database")

	db := DB{}
	db.SetConn(conn)

	return true, nil
}

// SetConn is method to deliver connection in struct
func (d DB) SetConn(conn *gorm.DB) {
	Conn = conn
}

// Get is method to get connection
func (d DB) Get() *gorm.DB {
	return Conn
}
