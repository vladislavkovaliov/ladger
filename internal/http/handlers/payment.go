package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
	service "github.com/vladislavkovaliov/ledger/internal/service"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a payment with amount and category
// @Tags payments
// @Accept json
// @Produce json
// @Param input body dto.CreatePaymentRequest true "Payment data"
// @Success 201 {object} dto.PaymentResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /payments [post]
func (h *PaymentHandler) Create(c *gin.Context) {
	var p payment.Payment

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.service.Create(c.Request.Context(), &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, p)
}

// ListPayments godoc
// @Summary List all payments
// @Description Get a list of all payments
// @Tags payments
// @Produce json
// @Success 200 {array} dto.PaymentResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /payments [get]
func (h *PaymentHandler) List(c *gin.Context) {
	payments, err := h.service.List(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, payments)
}
