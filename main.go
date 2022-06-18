package main

import (
	"log"
	"net/http"
	"strings"

	"momen/auth"
	"momen/handler"
	"momen/helper"
	"momen/transaction"
	"momen/users"

	"github.com/dgrijalva/jwt-go"
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


