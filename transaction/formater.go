package transaction

type TransactionFormater struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
}

func FormatTransaction(trans Transaction) TransactionFormater {
	transactionFormater := TransactionFormater{}

	transactionFormater.ID = trans.ID
	transactionFormater.UserID = trans.UserID
	transactionFormater.Name = trans.Name
	transactionFormater.Category = trans.Category
	transactionFormater.Description = trans.Description
	transactionFormater.Amount = trans.Amount

	return transactionFormater
}

func FormatTransactions(transactions []Transaction) []TransactionFormater {
	transactionsFormater := []TransactionFormater{}
	for _, trans := range transactions {
		trasactionFormater := FormatTransaction(trans)
		transactionsFormater = append(transactionsFormater, trasactionFormater)
	}
	return transactionsFormater
}
