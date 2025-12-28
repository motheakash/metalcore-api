package router

import (
	"metalcore-api/internal/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	// Global middlewares
	r.Use(gin.Logger(), gin.Recovery())

	// Health check endpoint
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// API versioning
	v1 := r.Group("/api/v1")

	// Public routes
	// auth.RegisterRoutes(v1, db)

	// User routes (no auth yet)
	user.RegisterRoutes(v1, db)

	return r
}
