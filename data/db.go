package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB

	Employee *employeesRepo
}

func NewDB(dsn string) *DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("An error occurred while connecting to the database")
	}

	dataBase := DB{
		conn:     db,
		Employee: newEmployeesRepo(db),
	}

	db.AutoMigrate(&Employee{})

	return &dataBase
}
