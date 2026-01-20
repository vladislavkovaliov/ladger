package router

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/vladislavkovaliov/ledger/internal/http/handlers"
)

func RegisterRouter(r *gin.Engine, payment *handlers.PaymentHandler, user *handlers.UserHandler) {
	r.GET("/health", handlers.HealthHandler)

	r.POST("/payments", payment.Create)
	r.GET("/payments", payment.List)

	r.GET("/users", user.List)
	r.POST("/users/create", user.Create)

}
