package handler

import (
	"fmt"
	"momen/helper"
	"momen/transaction"
	"momen/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetTransactions(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID
	transactions, err := h.service.GetTransactions(userID)

	metaError := helper.Meta{
		Message: "error to get transactions", Code: http.StatusBadRequest, Status: "error",
	}
	if err != nil {
		response := helper.APIResponse(metaError, nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	meta := helper.Meta{
		Message: "List of transactions", Code: http.StatusOK, Status: "success",
	}
	response := helper.APIResponse(meta, transaction.FormatTransactions(transactions))

	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.TransactionInput

	err := c.ShouldBindJSON(&input)
	metaError := helper.Meta{
		Message: "Failed to create transactions", Code: http.StatusBadRequest, Status: "error",
	}
	if err != nil {
		errs := ErrorValidationHandler(err)

		errorMessage := gin.H{"error": errs}
		response := helper.APIResponse(metaError, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	input.UserID = currentUser.ID
	newTransaction, err := h.service.CreateTransaction(input)
	fmt.Println(newTransaction)
	if err != nil {

		response := helper.APIResponse(metaError, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	meta := helper.Meta{
		Message: "Success to create transaction", Code: http.StatusOK, Status: "success",
	}
	response := helper.APIResponse(meta, transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler)UpdateTransaction(c *gin.Context)  {
	var inputID transaction.GetTransactionInputID

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		metaUnprocess := helper.Meta{
			Message: "Failed to update transactions", Code: http.StatusUnprocessableEntity, Status: "error",
		}
		errs := ErrorValidationHandler(err)

		errorMessage := gin.H{"error": errs}
		response := helper.APIResponse(metaUnprocess, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData transaction.TransactionInput
	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		metaUnprocess := helper.Meta{
			Message: "Failed to update transactions", Code: http.StatusUnprocessableEntity, Status: "error",
		}
		errs := ErrorValidationHandler(err)

		errorMessage := gin.H{"error": errs}
		response := helper.APIResponse(metaUnprocess, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	inputData.UserID = currentUser.ID

	updateTransaction, err := h.service.UpdateTransaction(inputID, inputData)

	if err != nil {
		metaBadRequest := helper.Meta{
			Message: "Failed to update transactions", Code: http.StatusBadRequest, Status: "error",
		}
		response := helper.APIResponse(metaBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	metaSuccess := helper.Meta{
		Message: "Success to update transaction", Code: http.StatusOK, Status: "success",
	}
	response := helper.APIResponse(metaSuccess, transaction.FormatTransaction(updateTransaction))
	c.JSON(http.StatusOK, response)
}