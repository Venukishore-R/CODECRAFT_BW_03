package handlers

import (
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/auth"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

type UserHandlerInterface interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
		})
		return
	}

	status, err := h.UserService.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error":   err.Error(),
			"message": "failed to create user",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "user created successfully",
		"user":    user,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}

	var loginReq *LoginReq
	if err := ctx.ShouldBindBodyWithJSON(&loginReq); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
		})
		return
	}

	status, token, err := h.UserService.Login(loginReq.Email, loginReq.Pass)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error":   err.Error(),
			"token":   nil,
			"message": "failed to login",
		})
		return
	}

	ctx.JSON(status, gin.H{
		"token":   token,
		"message": "login successful",
	})
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	claims, _ := ctx.Get("user")

	status, userDetails, err := h.UserService.UserProfile(claims.(*auth.Claims).Email)
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error":   err.Error(),
			"message": "failed to get user profile",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user": userDetails,
	})
}
