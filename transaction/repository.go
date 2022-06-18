package transaction

import "gorm.io/gorm"

type Repository interface {
	FindTransByID(userID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	FindByID(userID int) (Transaction, error)
	UpdateTransaction(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransByID(userID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("ID = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, err
}

func (r *repository) UpdateTransaction(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
