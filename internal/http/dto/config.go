package dto

type ConfigResponse struct {
	Status      string `json:"status" example:"ok"`
	Port        string `json:"port" example:"ok"`
	DatabaseUrl string `json:"databaseUrl" example:"ok"`
	Secret      string `json:"secret" example:"ok"`
}
