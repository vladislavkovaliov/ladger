package dto

type HealthResponse struct {
	Status string `json:"Status" example:"ok" binding:"required"`
}
