package handler

import (
	"fmt"
	"momen/auth"
	"momen/entities"
	"momen/helper"
	"momen/input_post"
	"momen/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.Service
	authService auth.AuthService
}

func NewUserHandler(userService services.Service, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegiterUser(c *gin.Context) {
	var input inputpost.RegisterInput
	err := c.ShouldBindJSON(&input)

	metaError := helper.Meta{
		Message: "Register Account Failed", Code: http.StatusUnprocessableEntity, Status: "Error",
	}

	if err != nil {
		errors := ErrorValidationHandler(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse(metaError, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)

	if err != nil {
		response := helper.APIResponse(metaError, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := helper.FormatUser(user, token)
	meta := helper.Meta{
		Message: "Account has been created", Code: http.StatusOK, Status: "Success",
	}

	response := helper.APIResponse(meta, formater)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input inputpost.LoginInput

	err := c.ShouldBindJSON(&input)
	metaError := helper.Meta{
		Message: "Login Failed", Code: http.StatusUnprocessableEntity, Status: "Error",
	}
	if err != nil {
		errors := ErrorValidationHandler(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	user, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)

	if err != nil {
		response := helper.APIResponse(metaError, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := helper.FormatUser(user, token)
	meta := helper.Meta{
		Message: "Login Successfully", Code: http.StatusOK, Status: "Success",
	}

	response := helper.APIResponse(meta, formater)

	c.JSON(http.StatusOK, response)
}

// cek email availabelity
func (h *userHandler) CheckEamilAvailablelity(c *gin.Context) {
	var input inputpost.CheckEamilInput

	err := c.ShouldBindJSON(&input)
	metaError := helper.Meta{
		Message: "Email checking failed", Code: http.StatusUnprocessableEntity, Status: "Error",
	}
	if err != nil {
		errors := ErrorValidationHandler(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been used"

	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	meta := helper.Meta{
		Message: metaMessage, Code: http.StatusOK, Status: "Success",
	}

	response := helper.APIResponse(meta, data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	data := gin.H{"is_uplaoded": false}
	metaError := helper.Meta{
		Message: "Failed to upload", Code: http.StatusBadRequest, Status: "error",
	}
	if err != nil {
		response := helper.APIResponse(metaError, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entities.User)

	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		response := helper.APIResponse(metaError, data)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		response := helper.APIResponse(metaError, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data = gin.H{"is_uplaoded": true}
	meta := helper.Meta{
		Message: "Avatar successfully to upload", Code: http.StatusOK, Status: "success",
	}
	response := helper.APIResponse(meta, data)

	c.JSON(http.StatusOK, response)

}
