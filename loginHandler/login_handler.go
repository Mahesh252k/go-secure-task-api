package loginhandler

import (
	"CRUD_API_PROJ/models"
	serviceauth "CRUD_API_PROJ/serviceAuth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	AuthService serviceauth.AuthService
}

func NewLoginHandler(AuthService serviceauth.AuthService) *LoginHandler {
	return &LoginHandler{
		AuthService: AuthService,
	}
}

func (h *LoginHandler) Register(c *gin.Context) {
	var registerRequest models.User
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthService.Register(registerRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginRequest models.User
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Login(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
