package main

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	log.Println("Running the db migrations...")
	dsn := "host=localhost user=go password=go dbname=golanghex port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}

	err = conn.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Error doing the migration", err)
	}
	log.Println("Finished", err)
}

func main() {
	Migrate()
}
