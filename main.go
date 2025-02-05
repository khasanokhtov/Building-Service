package main

import (
	"building-service/config"
	"building-service/database"
	_ "building-service/docs" //документация
	"building-service/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Building API
// @version 1.0
// @description API для управления строениями
// @host localhost:8080
// @BasePath /


func main(){
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()

	r.POST("/buildings", handlers.CreateBuildingHandler)
	r.GET("/buildings", handlers.GetBuildingsHandler)

	// Сваггер юзерИ
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}