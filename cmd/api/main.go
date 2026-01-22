package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/config"
	router "github.com/vladislavkovaliov/ledger/internal/http"
	handlers "github.com/vladislavkovaliov/ledger/internal/http/handlers"
	mongoClient "github.com/vladislavkovaliov/ledger/internal/infra/mongo"
	repository_payment "github.com/vladislavkovaliov/ledger/internal/repository/payment"
	repository_user "github.com/vladislavkovaliov/ledger/internal/repository/user"
	service "github.com/vladislavkovaliov/ledger/internal/service"

	// Swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/vladislavkovaliov/ledger/docs"
)

// @title Ledger API
// @version 1.0
// @description Personal finance ledger API
// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()

	fmt.Println("Starting server on port", cfg.Port)

	client, db, err := mongoClient.NewMongoClient(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	repoPayment := repository_payment.NewPaymentRepository(db.Collection("payments"))
	servicePayment := service.NewPaymentService(repoPayment)
	handlerPayment := handlers.NewPaymentHandler(servicePayment)

	repoUser := repository_user.NewUserRepository(db.Collection("users"))
	serviceUser := service.NewUserService(repoUser)
	handlerUser := handlers.NewUserHandler(serviceUser, *cfg)

	r := gin.Default()

	router.RegisterRouter(r, handlerPayment, handlerUser, cfg)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:" + cfg.Port)
}
