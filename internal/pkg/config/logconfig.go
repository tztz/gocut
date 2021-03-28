package config

import log "github.com/sirupsen/logrus"

// InitLogConfig configures logging
func InitLogConfig() {
	if IsProdProfileEnabled() {
		// Log as JSON
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// Log as plain text (default ASCII formatter)
		log.SetFormatter(&log.TextFormatter{})
	}
}
