package config

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func assertProp(t *testing.T, configsPath string, value int) {
	t.Helper()

	viper.Reset()
	err := InitAppConfigForPath(configsPath)
	assert.Nil(t, err, "should have initialized app config successfully but did not work")
	assert.Equal(t, value, viper.Get("foo.bar"), "could not find app property 'foo.bar'")
}

func assertNotFound(t *testing.T, path string) {
	t.Helper()

	viper.Reset()
	err := InitAppConfigForPath(path)
	var e viper.ConfigFileNotFoundError
	assert.True(t, errors.As(err, &e), "should not have found app config but it exists")
	assert.True(t, strings.HasPrefix(err.Error(), "app config file not found for profile"))
}

func TestInitAppConfig(t *testing.T) {
	assert.Equal(t, InitAppConfig(), InitAppConfigForPath(getConfigsPath()))
}

func TestAppConfigFileShouldBeReadAndHaveProp(t *testing.T) {
	configsPath := getProjectRoot() + "/test/configs"

	os.Setenv("RUN_PROFILE", "test")
	assertProp(t, configsPath, 815)

	os.Setenv("RUN_PROFILE", "dev")
	assertProp(t, configsPath, 5551234)

	os.Setenv("RUN_PROFILE", "prod")
	assertProp(t, configsPath, 4711)

	os.Setenv("RUN_PROFILE", "")
	assertProp(t, configsPath, 4711)
}

func TestViperShouldReturnErrorIfAppConfigFileNotFound(t *testing.T) {
	os.Setenv("RUN_PROFILE", "test")
	assertNotFound(t, "/tmp/foo-does-not-exist")

	os.Setenv("RUN_PROFILE", "foo-does-not-exist")
	assertNotFound(t, "/tmp/foo-does-not-exist")

	os.Setenv("RUN_PROFILE", "foo-does-not-exist")
	assertNotFound(t, "")
}

func TestViperShouldReturnErrorIfAppConfigFileFoundButCorrupt(t *testing.T) {
	configsPath := getProjectRoot() + "/test/configs"

	viper.Reset()
	os.Setenv("RUN_PROFILE", "corrupt")
	err := InitAppConfigForPath(configsPath)
	var e viper.ConfigParseError
	assert.True(t, errors.As(err, &e), "should have found corrupt app config but it is valid")
	assert.True(t, strings.HasPrefix(err.Error(), "app config file found but corrupt for profile [corrupt]: "))
}

func TestIsFooProfileEnabledFuncs(t *testing.T) {
	os.Setenv("RUN_PROFILE", "test")
	assert.True(t, IsTestProfileEnabled())

	os.Setenv("RUN_PROFILE", "dev")
	assert.True(t, IsDevProfileEnabled())

	os.Setenv("RUN_PROFILE", "prod")
	assert.True(t, IsProdProfileEnabled())

	os.Setenv("RUN_PROFILE", "")
	assert.True(t, IsProdProfileEnabled())

	os.Setenv("RUN_PROFILE", "foo-bar")
	assert.False(t, IsTestProfileEnabled())
	assert.False(t, IsDevProfileEnabled())
	assert.False(t, IsProdProfileEnabled())
}
