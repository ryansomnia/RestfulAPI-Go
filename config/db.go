package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	LocalHost     string
	LocalUsername string
	LocalPassword string
	LocalDatabase string
}

var db *sql.DB

func ConnectToDB() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file %v", err)
	}

	config := &Config{
		LocalHost:     os.Getenv("LOCAL_HOST"),
		LocalUsername: os.Getenv("LOCAL_USERNAME"),
		LocalPassword: os.Getenv("LOCAL_PASSWORD"),
		LocalDatabase: os.Getenv("LOCAL_DATABASE"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.LocalUsername,
		config.LocalPassword,
		config.LocalHost,
		config.LocalDatabase)

	d, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	db = d

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
