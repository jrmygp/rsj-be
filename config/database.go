package config

import (
	"fmt"
	"os"
	"server/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.UserRole{}, &models.User{}, &models.Shipper{}, &models.Customer{},
		&models.Port{}, &models.CostCharges{}, &models.Quotation{},
		&models.InvoiceImport{}, &models.InvoiceExport{}, &models.DoorToDoorInvoice{}, &models.SuratTugas{}, &models.Shipment{},
		&models.SuratJalan{}, &models.Warehouse{})

	return db
}
