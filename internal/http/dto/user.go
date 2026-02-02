package dto

type RegisterRequest struct {
	Email    string `json:"Email" example:"test@mail.com" binding:"required"`
	Password string `json:"Password" example:"123456" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"ID" example:"64f8c..." binding:"required"`
	Email     string `json:"Email" example:"test@mail.com" binding:"required"`
	CreatedAt string `json:"CreatedAt" example:"2026-01-20T00:00:00Z" binding:"required"`
}
