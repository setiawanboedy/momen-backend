package transaction

import "time"

type Transaction struct {
	ID          int    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID      int    `gorm:"size:36;index"`
	Description string `gorm:"size:100;not null"`
	Category    string `gorm:"size:100;not null"`
	Amount      int    `gorm:"size:100;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
