package transaction

import "errors"

type Service interface {
	GetTransactions(userID int) ([]Transaction, error)
	CreateTransaction(input TransactionInput) (Transaction, error)
	UpdateTransaction(inputID GetTransactionInputID, inputData TransactionInput)(Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactions(userID int) ([]Transaction, error) {
	transaction := []Transaction{}
	if userID != 0 {
		transactions, err := s.repository.FindTransByID(userID)
		if err != nil {
			return transactions, err
		}

		return transactions, nil
	}

	return transaction, errors.New("no trans found")

}

func (s *service) CreateTransaction(input TransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Name = input.Name
	transaction.Description = input.Description
	transaction.Category = input.Category
	transaction.Amount = input.Amount
	transaction.UserID = input.UserID
	newTransaction, err := s.repository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service)UpdateTransaction(inputID GetTransactionInputID, inputData TransactionInput)(Transaction, error)  {
	transaction, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return transaction, err
	}

	if transaction.UserID != inputData.UserID {
		return transaction, errors.New("not owner of transaction")
	}

	transaction.Name = inputData.Name
	transaction.Description = inputData.Description
	transaction.Category = inputData.Category
	transaction.Amount = inputData.Amount

	updateTransaction, err := s.repository.UpdateTransaction(transaction)
	if err != nil {
		return updateTransaction, err
	}
	return updateTransaction, nil
}
