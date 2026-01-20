package payment

import "time"

type Payment struct {
	ID         string
	Amount     int64
	CreateDate time.Time
	UpdateDate time.Time
	CategoryID string
}
