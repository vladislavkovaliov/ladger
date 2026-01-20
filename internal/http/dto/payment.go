package dto

type CreatePaymentRequest struct {
	Amount     int64  `json:"amount" example:"1500"`
	CategoryID string `json:"category_id" example:"cat_123"`
}

type PaymentResponse struct {
	ID         string `json:"id" example:"64f8c..."`
	Amount     int64  `json:"amount" example:"1500"`
	CategoryID string `json:"category_id" example:"cat_123"`
	CreateDate string `json:"create_date" example:"2026-01-20T00:00:00Z"`
	UpdateDate string `json:"update_date" example:"2026-01-20T00:00:00Z"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"internal server error"`
}
