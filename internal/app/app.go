package app

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"srv.tztz.io/example/gocut/internal/pkg/config"
	"srv.tztz.io/example/gocut/internal/pkg/middleware"
)

// Start is the entrypoint of the gocut service.
// Here everything is wired together.
func Start() {
	initConfigs()

	log.Info("Ahoi! This is gocut running with profile '" + config.GetRunProfile() + "'")

	r := gin.Default()
	r.LoadHTMLGlob(path.Join(config.GetProjectRoot(), "web/templates/*"))

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

	// listen and serve
	if err := r.Run(":" + viper.GetString("gin.port")); err != nil {
		log.Fatal("Could not start server. ", err)
	}
}

func initConfigs() {
	config.InitLogConfig()
	if err := config.InitAppConfig(); err != nil {
		log.Fatalf("Could not initialize app config: %s", err)
	}
	config.InitGinConfig()
}
