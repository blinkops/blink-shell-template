package main

import (
	"github.com/blinkops/plugin-sdk"
	"github.com/blinkops/plugin-sdk/plugin/config"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	executable, err := os.Executable()
	if err != nil {
		logrus.Error("Failed to get current executable with error: ", err)
		panic(err)
	}

	currentDirectory := filepath.Dir(executable)

	// Initialize the configuration.
	err = os.Setenv(config.ConfigurationPathEnvVar, "config.yaml")
	if err != nil {
		logrus.Error("Failed to set configuration env variable: ", err)
		panic(err)
	}

	plugin, err := NewShellRunner(currentDirectory)
	if err != nil {
		logrus.Error("Failed to create plugin implementation: ", err)
		panic(err)
	}

	err = plugin_sdk.Start(&*plugin)
}
