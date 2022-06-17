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

	// user, err := userService.LoginUser()


	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	//endpoint
	api.POST("/users", userHandler.RegiterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_chekers", userHandler.CheckEamilAvailablelity)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

}

//controller handler

