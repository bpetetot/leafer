package server

import (
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/server/api"
)

// Start the web server
func Start() {
	port := os.Getenv("PORT")
	router := gin.Default()

	router.Use(db.Middleware)
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	api.Routes("/api", router)

	router.Run(":" + port)
}
