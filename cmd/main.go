package main

import (
	"github.com/gin-gonic/gin"
	prometheus "srv.tztz.io/example/gocut/internal/pkg/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/metrics", prometheus.Handler())

	// listen and serve on port 3000
	r.Run(":3000")
}
