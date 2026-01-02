package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(rg *gin.RouterGroup, db *pgxpool.Pool) {
	// Initialize dependencies (Dependency Injection)
	repo := NewUserRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	// Register routes
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/:id", handler.GetByID)
		userGroup.GET("/", handler.GetAll)
		userGroup.POST("/", handler.Create)
		// userGroup.PUT("/:id", handler.Update)
		// userGroup.DELETE("/:id", handler.Delete)
	}
}
