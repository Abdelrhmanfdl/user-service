package handlers

import (
	"net/http"

	"github.com/Abdelrhmanfdl/user-service/internal/errs"
	"github.com/Abdelrhmanfdl/user-service/internal/models"
	"github.com/Abdelrhmanfdl/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	userService *service.UserService
}

func NewRouterHandler(userService *service.UserService) *RouterHandler {
	return &RouterHandler{
		userService: userService,
	}
}

func (h *RouterHandler) HandleLogin(ctx *gin.Context) {
	var loginBody models.DtoLoginRequest

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if token, err := h.userService.LoginUser(loginBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (h *RouterHandler) HandleSignup(ctx *gin.Context) {
	var signupBody models.DtoSignupRequest

	if err := ctx.ShouldBindJSON(&signupBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if token, err := h.userService.SignupUser(signupBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (h *RouterHandler) HandleGetUserData(ctx *gin.Context) {
	userId, isExisting := ctx.Params.Get("userId")

	if !isExisting {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs.NotFoundUser{}.Message})
	}

	if user, err := h.userService.GetUserData(userId); err != nil {
		// TODO: check error type
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	}
}
