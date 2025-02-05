package handlers

import (
	"building-service/models"
	"building-service/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBuildingHandler добавляет новое строение в БД
// @Summary Добавить строение
// @Description Добавляет новое строение в базу данных
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body models.Building true "Данные строения"
// @Success 201 {object} models.Building
// @Failure 400 {string} string "Ошибка валидации"
// @Router /buildings [post]
func CreateBuildingHandler(c *gin.Context){
	var building models.Building
	
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateBuilding(&building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Ошибка при сохранении"})
		return

	}

	c.JSON(http.StatusCreated, building)
}



// GetBuildingsHandler возвращает список строений с фильтрацией
// @Summary Получить список строений
// @Description Получает список строений с возможностью фильтрации по городу, году (диапазон) и количеству этажей
// @Tags buildings
// @Accept json
// @Produce json
// @Param city query string false "Город"
// @Param year_from query int false "Год (от)"
// @Param year_to query int false "Год (до)"
// @Param floors query int false "Количество этажей"
// @Success 200 {array} models.Building
// @Router /buildings [get]
func GetBuildingsHandler(c *gin.Context){

	city := c.Query("city")
	year, _ := strconv.Atoi(c.Query("year"))
	floors, _ := strconv.Atoi(c.Query("floors"))

	buildings, err := repository.GetBuildings(city, year, floors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки данных"})
		return
	}

	c.JSON(http.StatusOK, buildings)
}