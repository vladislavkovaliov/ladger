package payment

import "time"

type Payment struct {
	ID         string
	Amount     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CategoryID string
}
