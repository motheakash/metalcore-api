package user

import (
	"metalcore-api/internal/common"
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
		"data": ToUserResponse(user),
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	var pagination common.PaginationRequest
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Invalid pagination parameters",
			Message: err.Error(),
		})
		return
	}
	page := pagination.GetPage()
	page_size := pagination.GetPageSize()
	users, total, err := h.service.GetAll(c.Request.Context(), page, page_size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	user_response := ToUserListResponse(users)

	pagination_meta := common.NewPaginationMetadata(&pagination, total)

	c.JSON(http.StatusOK, common.PaginatedResponse{
		Data:       user_response,
		Pagination: pagination_meta,
	})
}

func (h *Handler) Create(c *gin.Context) {
	var payload CreateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validationErrors := common.FormatValidationErrors(err)
		c.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Validation failed",
			Message: "Please check the input fields",
			Details: validationErrors,
		})
		return
	}

	response, err := h.service.Create(c, payload)
	if err != nil {
		switch err {
		case ErrUsernameExists:
			c.JSON(http.StatusConflict, common.ErrorResponse{
				Status:  http.StatusConflict,
				Error:   "Username already exists",
				Message: "Please choose a different username",
			})
		default:
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Error:   "Internal server error",
				Message: "An unexpected error occurred",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been created successfully.",
		"data":    response,
	})

}
