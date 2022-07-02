package transaction

type TransactionInput struct {
	UserID      int
	ID          int    `json:"id" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
}

type GetTransactionInputID struct {
	ID int `uri:"id" binding:"required"`
}

type GetDetailTransactionID struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteTransactionInputID struct {
	ID int `uri:"id" binding:"required"`
}
