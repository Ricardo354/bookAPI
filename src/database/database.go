package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {

	dsn := "host=127.0.0.1 user=usuario password=1234 dbname=olist_api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}
	
	fmt.Println("Conexão bem sucedida")
	DBConn = db
}
