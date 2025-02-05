package repository

import (
	"building-service/database"
	"building-service/models"
)

func CreateBuilding(building *models.Building) error {
	return database.DB.Create(building).Error
}

func GetBuildings(city string, year int, floors int)([]models.Building, error){
	var buildings []models.Building
	query := database.DB

	if city != ""{
		query = query.Where("city = ?", city)
	}
	if year != 0 {
		query = query.Where("year = ?", year)

	}
	if floors != 0 {
		query = query.Where("floors = ?", floors)

	}

	err := query.Find(&buildings).Error
	return buildings, err
}