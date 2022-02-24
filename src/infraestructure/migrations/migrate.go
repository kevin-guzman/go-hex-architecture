package main

import (
	"golang-gingonic-hex-architecture/src/infraestructure"
	userEntity "golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	log.Println("Running the db migrations...")

	run_env := os.Getenv("ENVIRONMENT")
	path_env := "./env/development.env"
	switch run_env {
	case "test":
		path_env = "./env/testing.env"
	case "production":
		path_env = "./env/production.env"
	default:
		path_env = "./env/development.env"
	}

	if err := godotenv.Load(path_env); err != nil {
		log.Fatal("Error reading .env file\n", err)
	}

	dbProperties := infraestructure.DatabaseConnectionProperties{
		DATABASE_TYPE:     os.Getenv("DATABASE_TYPE"),
		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
	}
	DATABSE_STRING_CONNECTION := infraestructure.CreateDatabaseStringConnetion(&dbProperties)
	conn, err := gorm.Open(postgres.Open(DATABSE_STRING_CONNECTION), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}

	entities := map[string]interface{}{
		"user": &userEntity.User{},
	}

	for name, entity := range entities {
		err = conn.AutoMigrate(entity)
		if err != nil {
			log.Fatal("Error migrating the: ", name, "entity ", err)
			break
		}
	}

	log.Println("Finished")
}

func main() {
	Migrate()
}
