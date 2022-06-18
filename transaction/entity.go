package transaction

import "time"

type Transaction struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Category    string
	Amount      int
	CreatedAt   time.Time
	UpdatedAt    time.Time
}
