package handler

import (
	"momen/helper"
	"momen/transaction"
	"momen/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) * transactionHandler  {
	return &transactionHandler{service}
}

func (h *transactionHandler)GetTransactions(c *gin.Context)  {

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
	response := helper.APIResponse(meta, transactions)

	c.JSON(http.StatusOK, response)

}