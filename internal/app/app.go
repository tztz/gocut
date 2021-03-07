package app

import (
	"github.com/gin-gonic/gin"
	healthcheck "srv.tztz.io/example/gocut/internal/pkg/middleware"
	prometheus "srv.tztz.io/example/gocut/internal/pkg/middleware"
)

// Start is the entrypoint of the gocut service.
// Here everything is wired together.
func Start() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/admin/metrics", prometheus.PrometheusHandler())

	r.GET("/admin/healthcheck", healthcheck.HealthcheckHandler())

	// listen and serve on port 3000
	r.Run(":3000")
}
