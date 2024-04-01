package db

import (
	"gorm-concept/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db gorm.DB
}

func New() Database {
	db, err := gorm.Open(sqlite.Open("gorm-concept.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.Product{})

	return Database{
		Db: *db,
	}


}
