package db

import (
	"github.com/gin-gonic/gin"
)

// Middleware initiate database and give db connection to routes
func Middleware(c *gin.Context) {
	db := Setup()

	c.Set("db", db)
	c.Next()
}
