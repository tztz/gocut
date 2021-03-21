package config

import (
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitAppConfig sets up Viper.
func InitAppConfig() {
	// load base config:
	mergeInViperConfig("")
	// load profile-specific config and override base config:
	switch os.Getenv("PROFILE") {
	case "test":
		mergeInViperConfig("test")
	case "dev":
		mergeInViperConfig("dev")
	}
}

// ReadAppConfig reads the application config file and sets configuration defaults.
func ReadAppConfig() error {
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Application config file not found")
		} else {
			log.Error("Application config file found but corrupt. ", err)
		}
	} else {
		log.Info("Application config file successfully read")
	}

	return err
}

// mergeInViperConfig loads and merges the application config for the given profile.
func mergeInViperConfig(profile string) {
	appendix := ""
	if profile != "" {
		appendix = "_" + profile
	}
	viper.SetConfigName("application" + appendix)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getConfigsPath())
	viper.MergeInConfig()
}

// getProjectRoot returns the absolute path to the project root.
func getProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logrus.Fatal("Unable to retrieve runtime caller information")
	}
	// "filename" is an absolute path to this file; move up 3 levels to reach <project-root>
	return path.Clean(path.Join(path.Dir(filename), "..", "..", ".."))
}

// getConfigsPath returns the absolute path to the <project-root>/configs directory.
func getConfigsPath() string {
	return path.Join(getProjectRoot(), "configs")
}
