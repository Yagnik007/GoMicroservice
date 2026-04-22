package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myorg/myservice/internal/handlers"
	"github.com/myorg/myservice/pkg/response"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes sets up all API endpoints
func RegisterRoutes(router *gin.Engine, itemHandler *handlers.ItemHandler) {

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "Service is healthy", nil)
	})

	api := router.Group("/api/v1")
	{
		items := api.Group("/items")
		{
			items.GET("", itemHandler.GetItems)
			items.GET("/:id", itemHandler.GetItem)
			items.POST("", itemHandler.CreateItem)
		}
	}
}
