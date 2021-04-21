package main

import (
	"github.com/blinkops/plugin-sdk"
	"github.com/blinkops/plugin-sdk/plugin/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	// Get the current directory.
	currentDirectory, err := os.Getwd()
	if err != nil {
		logrus.Error("Failed getting current directory: ", err)
		panic(err)
	}

	// Initialize the configuration.
	err = os.Setenv(config.ConfigurationPathEnvVar, path.Join(currentDirectory, "config.yaml"))
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
