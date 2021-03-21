package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAppConfigFileShouldBeReadAndHaveProp(t *testing.T) {
	// always start fresh:
	viper.Reset()

	InitAppConfig()
	err := ReadAppConfig()

	assert.Nil(t, err, "should have read app config successfully but did not work")

	assert.Equal(t, 815, viper.Get("foo.bar"), "could not find app property 'foo.bar'")
}

func TestViperShouldReturnErrorIfAppConfigFileNotFound(t *testing.T) {
	// always start fresh:
	viper.Reset()

	err := ReadAppConfig()

	_, isExpectedError := err.(viper.ConfigFileNotFoundError)
	assert.True(t, isExpectedError, "should not have found app config but it exists")
}
