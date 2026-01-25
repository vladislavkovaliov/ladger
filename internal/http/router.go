package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/config"
	handlers "github.com/vladislavkovaliov/ledger/internal/http/handlers"
	"github.com/vladislavkovaliov/ledger/internal/middleware"
)

func RegisterRouter(r *gin.Engine, payment *handlers.PaymentHandler, user *handlers.UserHandler, cfg *config.Config) {
	r.GET("/health", handlers.HealthHandler)

	r.GET("/config", func(c *gin.Context) {
		handlers.ConfigHandler(c, cfg)
	})

	r.POST("/auth/create", user.Create)

	r.POST("/auth/login", user.Login)

	auth := r.Group("/")

	auth.Use(middleware.JWTAuth(cfg.Secret))
	{
		auth.GET("/payments", payment.List)
		auth.POST("/payments", payment.Create)
		auth.GET("/users", user.List)
	}
}
