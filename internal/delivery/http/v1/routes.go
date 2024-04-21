package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/commands-list")
		api.GET("/command/:id")
		api.POST("/command")
		api.POST("/command/:id")
		api.GET("/command")
		api.DELETE("command/:id")

	}
	return router
}
