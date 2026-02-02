package dto

type CreatePaymentRequest struct {
	Amount     int64  `json:"Amount" example:"1500" binding:"required"`
	CategoryID string `json:"CategoryID" example:"cat_123" binding:"required"`
}

type PaymentResponse struct {
	ID         string `json:"ID" example:"64f8c..." binding:"required"`
	Amount     int64  `json:"Amount" example:"1500" binding:"required"`
	CategoryID string `json:"CategoryID" example:"cat_123" binding:"required"`
	CreateDate string `json:"CreateDate" example:"2026-01-20T00:00:00Z" binding:"required"`
	UpdateDate string `json:"UpdateDate" example:"2026-01-20T00:00:00Z" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"internal server error" binding:"required"`
}
