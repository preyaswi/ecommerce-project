package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"newone/entity"
	"newone/models"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Dsn string
var KEY5 = "host=localhost user=preya dbname=preya password=preyaswi port=5432 sslmode=disable"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Dsn = os.Getenv("KEY5")

}

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(KEY5), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	Db = db
	Db.AutoMigrate(&models.Signup{}, &entity.OtpKey{}, &entity.User{})
	return db, nil
}

func ConnectToTestDB() (*sql.DB, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
