package main

import (
	"github.com/blinkops/blink-plugin-base/plugin"
	"github.com/blinkops/plugin-sdk"
	"github.com/blinkops/plugin-sdk/plugin/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {

	log.SetLevel(log.DebugLevel)

	executable, err := os.Executable()
	if err != nil {
		log.Error("Failed to get current executable with error: ", err)
		panic(err)
	}

	currentDirectory := filepath.Dir(executable)

	// Initialize the configuration.
	err = os.Setenv(config.ConfigurationPathEnvVar, "config.yaml")
	if err != nil {
		log.Error("Failed to set configuration env variable: ", err)
		panic(err)
	}

	plugin, err := plugin.NewShellRunner(currentDirectory)
	if err != nil {
		log.Error("Failed to create plugin implementation: ", err)
		panic(err)
	}

	err = plugin_sdk.Start(&*plugin)
}

