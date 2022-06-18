package transaction

import "errors"

type Service interface {
	FindTrans(userID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindTrans(userID int) ([]Transaction, error) {
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