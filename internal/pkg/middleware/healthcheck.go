package middleware

import (
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

// HealthcheckHandler attaches the healthcheck middleware
func HealthcheckHandler() gin.HandlerFunc {
	return healthcheck.Default()
}
