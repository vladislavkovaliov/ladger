package dto

type LoginRequest struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"Token" binding:"required"`
}
