package transaction

import "time"

type Transaction struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Amount      int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
