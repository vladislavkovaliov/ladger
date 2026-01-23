package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vladislavkovaliov/ledger/internal/config"
	router "github.com/vladislavkovaliov/ledger/internal/http"
	handlers "github.com/vladislavkovaliov/ledger/internal/http/handlers"
	mongoClient "github.com/vladislavkovaliov/ledger/internal/infra/mongo"
	repository_payment "github.com/vladislavkovaliov/ledger/internal/repository/payment"
	repository_user "github.com/vladislavkovaliov/ledger/internal/repository/user"
	service "github.com/vladislavkovaliov/ledger/internal/service"

	// Swagger
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/vladislavkovaliov/ledger/docs"
)

// @title Ledger API
// @version 1.0
// @description Personal finance ledger API
// @BasePath /
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

	r.Use(cors.Default())

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s - [%s] \"%s %s\" %d %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	router.RegisterRouter(r, handlerPayment, handlerUser, cfg)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(
	// 	swaggerFiles.Handler,
	// 	ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.Host, cfg.Port)),
	// ))

	// r.GET("/swagger/*any", func(c *gin.Context) {

	// 	url := "/swagger/doc.json"

	// 	c.Request.URL.Path = url
	// 	c.Writer.Header().Set("Content-Type", "application/json")
	// 	http.ServeFile(c.Writer, c.Request, "./docs/swagger.json")
	// })

	r.Run("0.0.0.0:" + cfg.Port)
}
