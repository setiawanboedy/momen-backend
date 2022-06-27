package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"momen/auth"
	"momen/handler"
	"momen/helper"
	"momen/migration"
	"momen/transaction"
	"momen/users"

	"momen/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// database
	_, dbConfig := utils.DatabaseSettings()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=enable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// cek database
	if err != nil {
		log.Fatal(err.Error())
	}

	// automigrate
	for _, model := range migration.MigrationModels(){
		err := db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}

	userRepository := users.NewRepository(db)
	transRepository := transaction.NewRepository(db)

	userService := users.NewService(userRepository)
	authService := auth.NewService()
	transService := transaction.NewService(transRepository)


	userHandler := handler.NewUserHandler(userService, authService)
	transHandler := handler.NewTransactionHandler(transService)

	router := gin.Default()
	router.Static("/images","./images")

	api := router.Group("/api/v1")

	//endpoint
	api.POST("/register", userHandler.RegiterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_chekers", userHandler.CheckEamilAvailablelity)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/transactions", authMiddleware(authService, userService), transHandler.GetTransactions)
	api.POST("/transaction", authMiddleware(authService, userService), transHandler.CreateTransaction)
	api.PUT("/transaction/:id", authMiddleware(authService, userService), transHandler.UpdateTransaction)
	api.DELETE("/transaction/:id", authMiddleware(authService, userService), transHandler.DeleteTransaction)

	router.Run()

}


func authMiddleware(authService auth.AuthService, userService users.Service) gin.HandlerFunc {
	// Middleware
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			meta := helper.Meta{
				Message: "Unauthorized", Code: http.StatusUnauthorized, Status: "error",
			}
			response := helper.APIResponse(meta, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			meta := helper.Meta{
				Message: "Unauthorized", Code: http.StatusUnauthorized, Status: "error",
			}
			response := helper.APIResponse(meta, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			meta := helper.Meta{
				Message: "Unauthorized", Code: http.StatusUnauthorized, Status: "error",
			}
			response := helper.APIResponse(meta, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			meta := helper.Meta{
				Message: "Unauthorized", Code: http.StatusUnauthorized, Status: "error",
			}
			response := helper.APIResponse(meta, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}


