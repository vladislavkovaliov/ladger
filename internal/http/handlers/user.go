package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/auth"
	"github.com/vladislavkovaliov/ledger/internal/config"
	"github.com/vladislavkovaliov/ledger/internal/domain/user"
	"github.com/vladislavkovaliov/ledger/internal/http/dto"
	"golang.org/x/crypto/bcrypt"

	service "github.com/vladislavkovaliov/ledger/internal/service"
)

type UserHandler struct {
	service *service.UserService
	jwt     config.Config
}

func NewUserHandler(s *service.UserService, jwtCfg config.Config) *UserHandler {
	return &UserHandler{
		service: s,
		jwt:     jwtCfg,
	}
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
// @Router /auth/create [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var u user.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	foundUser, err := h.service.FindByEmail(c.Request.Context(), req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if foundUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email " + foundUser.Email + " already exist."})

		return
	}

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	u.Email = req.Email
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
// @Tags users
// @Produce json
// @Security BearerAuth
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

// Login godoc
// @Summary Login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})

		return
	}

	u, err := h.service.FindByEmail(c.Request.Context(), req.Email)

	if err != nil {
		c.JSON(401, dto.ErrorResponse{Error: "invalid credentials"})

		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		fmt.Println(err)
		c.JSON(401, dto.ErrorResponse{Error: "invalid credentials"})

		return
	}

	token, err := auth.GenerateToken(
		u.ID,
		h.jwt.Secret,
		h.jwt.Expiration,
	)
	if err != nil {
		c.JSON(500, dto.ErrorResponse{Error: err.Error()})

		return
	}

	c.JSON(200, dto.LoginResponse{Token: token})
}
