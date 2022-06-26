package migration

import (
	"momen/transaction"
	"momen/users"
)

type Model struct {
	Model interface{}
}

func MigrationModels() []Model {
	models := []Model{
		{Model: users.User{}},
		{Model: transaction.Transaction{}},
	}
	return models
}