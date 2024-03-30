package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DataBase struct {
	Connection *sql.DB
}

func New() DataBase {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading env file", envErr)
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gin_note_app",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Could not connect to database --- ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Could not connect to database --- ", pingErr)
	}

	return DataBase{Connection: db}
}
