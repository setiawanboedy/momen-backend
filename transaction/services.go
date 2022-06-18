package transaction

import "errors"

type Service interface {
	GetTransactions(userID int) ([]Transaction, error)
	CreateTransaction(input TransactionInput) (Transaction, error)
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
