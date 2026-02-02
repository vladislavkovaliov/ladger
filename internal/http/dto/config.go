package dto

type ConfigResponse struct {
	Status      string `json:"Status" example:"ok" binding:"required"`
	Port        string `json:"Port" example:"8080" binding:"required"`
	DatabaseUrl string `json:"http://localhost:27017" example:"ok" binding:"required"`
	Secret      string `json:"Secret" example:"ok" binding:"required"`
}
