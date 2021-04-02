package config

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGinConfig() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	gin.SetMode(viper.GetString("gin.mode"))
}
