package config

import log "github.com/sirupsen/logrus"

// InitLogConfig configures logging
func InitLogConfig() {
	if IsProdProfileEnabled() {
		// Log as JSON instead of the default ASCII formatter
		log.SetFormatter(&log.JSONFormatter{})
	}
}
