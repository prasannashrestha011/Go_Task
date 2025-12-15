package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	db, err :=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		return err
	}
	sqlDB,err:=db.DB()
	if err!=nil{
		return err
	}
	if err:=sqlDB.Ping();err!=nil{
		log.Println("Database connection faild: ",err)
		return err
	}
	DB=db
	log.Println("SUCCESS: db connection successfull")
	return nil
}