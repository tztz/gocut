package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"srv.tztz.io/example/gocut/internal/pkg/config"
	"srv.tztz.io/example/gocut/internal/pkg/middleware"
)

// Start is the entrypoint of the gocut service.
// Here everything is wired together.
func Start() {
	config.InitLogConfig()
	config.InitAppConfig()

	log.Info("Ahoi! This is gocut running with profile '" + config.GetRunProfile() + "'")

	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "gocut - Ahoi!",
		})
	})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/admin/metrics", middleware.PrometheusHandler())

	r.GET("/admin/healthcheck", middleware.HealthcheckHandler())

	// listen and serve on port 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Could not start server. ", err)
	}
}
