package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	user, err := h.service.GetByID(c.Request.Context(), userID)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		case ErrUserInactive:
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User is inactive",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
