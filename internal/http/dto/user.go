package dto

type RegisterRequest struct {
	Email    string `json:"email" example:"test@mail.com"`
	Password string `json:"password" example:"123456"`
}

type UserResponse struct {
	ID         string `json:"id" example:"64f8c..."`
	Email      string `json:"email" example:"test@mail.com"`
	CreateDate string `json:"create_date" example:"2026-01-20T00:00:00Z"`
}
