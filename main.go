package main

import (
	"fmt"
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

	trans,_ := transService.FindTrans(20)
	fmt.Println(len(trans))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	//endpoint
	api.POST("/users", userHandler.RegiterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_chekers", userHandler.CheckEamilAvailablelity)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

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

// ambil nilai header authorization
// dari header ambil dari token
// validasi token
// ambil user_id
// ambil user dari db berdasarkan id lewat sevice
// set context isinya user
