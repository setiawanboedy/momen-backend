package handler

import (

	"momen/helper"
	"momen/input_post"
	"momen/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.Service
}

func NewUserHandler(userService services.Service) *userHandler {
	return &userHandler{userService}
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

	formater := helper.FormatUser(user, "token")
	meta := helper.Meta{
		Message: "Account has been created", Code: http.StatusOK, Status: "Success",
	}

	response := helper.APIResponse(meta, formater)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context)  {
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

	formater := helper.FormatUser(user, "token")
	meta := helper.Meta{
		Message: "Login Successfully", Code: http.StatusOK, Status: "Success",
	}

	response := helper.APIResponse(meta, formater)

	c.JSON(http.StatusOK, response)
}