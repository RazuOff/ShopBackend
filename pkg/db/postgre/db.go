package postgre

import (
	"fmt"
	"log"
	"os"

	"ginexample.com/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DeleteAllTables(); err != nil {
		log.Fatal(err.Error())
	}

	// Миграция схемы
	DB.AutoMigrate(&testProducts, &testUsers, &carts, &models.CartProduct{})

	insertTestData()
}

func insertTestData() {

	DB.Create(testProducts)
	DB.Create(testUsers)
	DB.Create(carts)
}

func DeleteAllTables() error {
	err := DB.Migrator().DropTable(&models.CartProduct{}, &models.Cart{}, &models.Product{}, &models.User{})
	return err
}
