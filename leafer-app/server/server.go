package server

import (
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/bpetetot/leafer/db"
)

// Start starts the web server
func Start() {
	port := os.Getenv("PORT")
	router := gin.Default()

	DB := db.Setup()
	defer DB.Close()

	router.Use(func(c *gin.Context) {
		c.Set("db", DB)
		c.Next()
	})

	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(static.Serve("/metadata", static.LocalFile("./.metadata", true)))

	api := router.Group("/api")

	library := NewLibraryHandlers(DB)
	api.GET("/libraries", library.Find)
	api.POST("/libraries", library.Create)
	api.GET("/libraries/:id", library.Get)
	api.DELETE("/libraries/:id", library.Delete)
	api.GET("/libraries/:id/scan", library.Scan)

	media := NewMediaHandlers(DB)
	api.GET("/media", media.Search)
	api.GET("/media/:id", media.Get)
	api.GET("/media/:id/content", media.GetContent)
	api.PATCH("/media/:id/read", media.MarkAsRead)
	api.PATCH("/media/:id/unread", media.MarkAsUnread)

	router.Run(":" + port)
}
