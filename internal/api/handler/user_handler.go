package handler

import (
	"dummies-backend/internal/api/dto"
	"dummies-backend/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")

	user, err := h.uc.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.UpdateProfile(c.Request.Context(), userID, req.UserHandle, req.UserName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userID := c.GetString("userID")

	if err := h.uc.DeleteAccount(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account deleted"})
}
