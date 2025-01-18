package handlers

import (
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Admin *services.Admin
}
type AdminHandlerInterface interface {
	GetUsers(ctx *gin.Context)
}

func NewAdminHandler(Admin *services.Admin) *AdminHandler {
	return &AdminHandler{
		Admin: Admin,
	}
}

func (h *AdminHandler) GetUsers(ctx *gin.Context) {
	statusCode, users, err := h.Admin.GetUsers()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error(), "message": "unable to get users"})
		return
	}

	ctx.JSON(statusCode, gin.H{"users": users})
}

func (h *AdminHandler) GetUser(ctx *gin.Context) {
	type Req struct {
		Email string `json:"email"`
	}

	var req *Req
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
		})

		return
	}

	code, user, err := h.Admin.GetUser(req.Email)
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error(), "message": "user not found"})
		return
	}

	ctx.JSON(code, gin.H{"user": user})
}

func (h *AdminHandler) Create(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
		})

		return
	}

	code, err := h.Admin.Create(user)
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error(), "message": "unable to create user"})
		return
	}

	ctx.JSON(code, gin.H{"message": "user created successfully"})
}

func (h *AdminHandler) Update(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error":  err.Error(),
			"messag": "invalid request payload",
		})
		return
	}

	code, result, err := h.Admin.UpdateUser(user, user.Email)
	if err != nil {
		ctx.JSON(code, gin.H{
			"error":   err.Error(),
			"message": "unable to update user",
		})
		return
	}

	ctx.JSON(code, gin.H{"user": result})
}

func (h *AdminHandler) Delete(ctx *gin.Context) {
	type Req struct {
		Email string `json:"email"`
	}

	var req *Req
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "invalid request payload",
		})
		return
	}

	code, err := h.Admin.Delete(req.Email)
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error(), "message": "unable to delete user"})
		return
	}

	ctx.JSON(code, gin.H{"message": "user deleted successfully"})
}
