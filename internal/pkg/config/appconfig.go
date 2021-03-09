package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitAppConfig reads the application config file and sets configuration defaults
func InitAppConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	// Depending on the directory in which the service has been started different config paths must be considered:
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Config file not found")
		} else {
			log.Error("Config file found but", err)
		}
	} else {
		log.Info("Config file successfully read")
		// TODO: remove this test line:
		log.Info("Example value: ", viper.Get("foo.bar"))
	}
}
