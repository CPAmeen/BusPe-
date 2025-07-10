package config

import (
	"fmt"
	"log"
	"os"

	"github.com/CPAmeen/buspe/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecttoDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("✅ Connected to PostgreSQL database remotely.")
}

//	func Syncdb() {

//			&models.Admins{},
//			&models.Owners{},
//			&models.Hotels{},
//			&models.Rooms{},
//			&models.Revenue{},
//			&models.Wallet{},
//		)
//		if err != nil {
//			fmt.Println("Couldn't Migrate tables", err)
//		} else {
//			fmt.Println("Successfully migrated", err)
//		}
//	}
func Sync() {
	// AutoMigrate all models // Assuming db is a global variable initialized elsewhere
	if DB == nil {
		panic("db instance is nil")
	}
	err := DB.AutoMigrate(
		&models.Users{},
		&models.Transaction{},
		&models.Owner{},
		&models.Owner{},
	)
	if err != nil {
		fmt.Println("Error while migration", err)
	} else {
		fmt.Println("Migration successful")
	}
}
