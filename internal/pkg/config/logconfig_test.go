package config

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInitLogConfig(t *testing.T) {
	os.Setenv("RUN_PROFILE", "prod")
	InitLogConfig()
	_, ok := log.StandardLogger().Formatter.(*log.JSONFormatter)
	assert.True(t, ok)

	os.Setenv("RUN_PROFILE", "")
	InitLogConfig()
	_, ok = log.StandardLogger().Formatter.(*log.JSONFormatter)
	assert.True(t, ok)

	os.Setenv("RUN_PROFILE", "dev")
	InitLogConfig()
	_, ok = log.StandardLogger().Formatter.(*log.TextFormatter)
	assert.True(t, ok)

	os.Setenv("RUN_PROFILE", "test")
	InitLogConfig()
	_, ok = log.StandardLogger().Formatter.(*log.TextFormatter)
	assert.True(t, ok)
}
