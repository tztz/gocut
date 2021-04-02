package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	prodProfile = "prod"
	testProfile = "test"
	devProfile  = "dev"
)

var errAppConfigFileNotFound = errors.New("app config file not found")
var errAppConfigFileFoundButCorrupt = errors.New("app config file found but corrupt")

// InitAppConfig sets up the Viper config.
func InitAppConfig() error {
	return InitAppConfigForPath(getConfigsPath())
}

// InitAppConfigForPath sets up the Viper config,
// searches in given path for config files.
func InitAppConfigForPath(configPath string) error {
	if configPath == "" {
		viper.AddConfigPath(getConfigsPath())
	} else {
		viper.AddConfigPath(configPath)
	}

	// load base (aka prod) config:
	if err := mergeInViperConfig(""); err != nil {
		return err
	}

	if profile := GetRunProfile(); profile != "" && profile != prodProfile {
		// load profile-specific config and override base/prod config:
		if err := mergeInViperConfig(profile); err != nil {
			return err
		}
	}

	return nil
}

func GetRunProfile() string {
	profile := os.Getenv("RUN_PROFILE")
	if profile == "" {
		return prodProfile
	}
	return profile
}

func IsTestProfileEnabled() bool {
	return GetRunProfile() == testProfile
}

func IsDevProfileEnabled() bool {
	return GetRunProfile() == devProfile
}

func IsProdProfileEnabled() bool {
	return GetRunProfile() == prodProfile
}

// GetProjectRoot returns the absolute path to the project root.
func GetProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to retrieve runtime caller information")
	}
	// "filename" is an absolute path to this file; move up 3 levels to reach <project-root>
	return path.Clean(path.Join(path.Dir(filename), "..", "..", ".."))
}

// mergeInViperConfig loads and merges the application config for the given profile.
func mergeInViperConfig(profile string) error {
	appendix := ""
	if profile != "" {
		appendix = "_" + profile
	}

	viper.SetConfigName("application" + appendix)
	viper.SetConfigType("yaml")

	if err := viper.MergeInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if errors.As(err, &e) {
			return fmt.Errorf("%v for profile [%s]: %w", errAppConfigFileNotFound, profile, err)
		}
		return fmt.Errorf("%v for profile [%s]: %w", errAppConfigFileFoundButCorrupt, profile, err)
	}

	log.Infof("App config file successfully read and merged for profile [%s]", profile)

	return nil
}

// getConfigsPath returns the absolute path to the <project-root>/configs directory.
func getConfigsPath() string {
	return path.Join(GetProjectRoot(), "configs")
}
