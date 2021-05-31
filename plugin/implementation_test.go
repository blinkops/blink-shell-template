package plugin

import (
	"fmt"
	"github.com/blinkops/blink-sdk/plugin"
	"github.com/blinkops/blink-sdk/plugin/config"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
	"testing"
)

type pluginBaseSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, &pluginBaseSuite{})
}

var plug *ShellRunner

func (s *pluginBaseSuite) SetupSuite() {
	err := os.Setenv(config.ConfigurationPathEnvVar, "config.yaml")
	if err != nil {
		assert.Fail(s.T(), "err", err)
	}
	plug, err = NewShellRunner("../test")
	if err != nil {
		assert.Fail(s.T(), "err", err)
	}
}

func (s *pluginBaseSuite) TestLoad() {
	assert.Equal(s.T(), "test123", plug.description.Name)
	assert.Equal(s.T(), "test plugin", plug.description.Description)
	assert.Equal(s.T(), "gotest", plug.description.Provider)
	assert.Equal(s.T(), []string{"hpux", "aix"}, plug.description.Tags)
	assert.Equal(s.T(), 1, len(plug.actions))
	bashAction := plug.actions[0]
	assert.Equal(s.T(), "bash", bashAction.Name)
	assert.Equal(s.T(), "bash_action.sh", bashAction.EntryPoint)
}

func (s *pluginBaseSuite) TestExecute() {
	response, err := plug.ExecuteAction(nil, &plugin.ExecuteActionRequest{
		Name: "bash",
		Parameters: map[string]string{
			"cmd": "pwd",
		},
	})
	assert.Nil(s.T(), err)
	fmt.Printf("response: %v\n", response)
	fmt.Println("result: ", string(response.Result))
	wd, err := os.Getwd()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), wd, strings.TrimSpace(string(response.Result)))
}

func (s *pluginBaseSuite) TestExecuteWithSpace() {
	response, err := plug.ExecuteAction(nil, &plugin.ExecuteActionRequest{
		Name: "bash",
		Parameters: map[string]string{
			"cmd": "echo hello world",
		},
	})
	if err != nil {
		log.Error(err)
	}
	assert.Nil(s.T(), err)
	fmt.Printf("response: %v\n", response)
	fmt.Println("result: ", string(response.Result))
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "hello world", strings.TrimSpace(string(response.Result)))
}