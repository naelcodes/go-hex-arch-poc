package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PostgresConnection() *gorm.DB {

	dsn := "host=ab.cjyodqyuof30.us-east-1.rds.amazonaws.com user=postgres  password=Neema_2023 dbname=airbooks port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("postgres database connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	return db

}
