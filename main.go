package main

import (
	"log"

	"momen/handler"
	"momen/repositories"
	"momen/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// database
	dsn := "root:@tcp(127.0.0.1:3306)/momen?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek database
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repositories.NewRepository(db)
	userService := services.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	//endpoint
	api.POST("/users", userHandler.RegiterUser)

	router.Run()

}

//controller handler

