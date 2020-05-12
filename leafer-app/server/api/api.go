package api

import (
	"github.com/gin-gonic/gin"
)

// Routes for API
func Routes(route string, router *gin.Engine) {
	api := router.Group(route)

	api.GET("/libraries", ListLibraries)
	api.POST("/libraries", CreateLibrary)
	api.GET("/libraries/:id", FindLibrary)
	api.PATCH("/libraries/:id", UpdateLibrary)
	api.DELETE("/libraries/:id", DeleteLibrary)
	api.GET("/libraries/:id/scan", ScanLibraryAsync)

	api.GET("/media", SearchMedia)
	api.GET("/media/:id", GetMedia)
	api.GET("/media/:id/content", GetMediaContent)
}
