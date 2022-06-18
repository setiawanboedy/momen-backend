package transaction

type TransactionInput struct {
	UserID      int
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
}

type GetTransactionInputID struct {
	ID int `uri:"id" binding:"required"`
}