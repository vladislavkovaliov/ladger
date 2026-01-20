package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/domain/user"
	"golang.org/x/crypto/bcrypt"

	service "github.com/vladislavkovaliov/ledger/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// RegisterUser godoc
// @Summary Register new user
// @Description Create a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "User registration data"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/create [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	foundUser, err := h.service.FindByEmail(c.Request.Context(), u.Email)

	if foundUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email " + foundUser.Email + " already exist."})

		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	u.PasswordHash = string(hash)

	if err := h.service.Create(c.Request.Context(), &u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, u)
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all users (emails only)
// @Tags auth
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users [get]
func (h *UserHandler) List(c *gin.Context) {
	users, err := h.service.List(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, users)
}
