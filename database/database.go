package database

import (
	"building-service/config"
	"building-service/models"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDB(){

	port, err := strconv.Atoi(config.AppConfig.Database.Port)
	if err != nil {
		log.Fatalf("Ошибка конверта порта в число: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DB)

	// var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	log.Println("Подключени к базе данных успешно!")


	DB.AutoMigrate(&models.Building{})
}