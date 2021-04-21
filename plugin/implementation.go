package main

import (
	"errors"
	"github.com/blinkops/plugin-sdk/plugin"
	actions2 "github.com/blinkops/plugin-sdk/plugin/actions"
	config2 "github.com/blinkops/plugin-sdk/plugin/config"
	description2 "github.com/blinkops/plugin-sdk/plugin/description"
	"github.com/sirupsen/logrus"
	"os/exec"
	"path"
)

type ShellRunner struct {
	actions     []plugin.Action
	description plugin.Description
}

func (p *ShellRunner) Describe() plugin.Description {
	logrus.Debug("Handling Describe request!")
	return p.description
}

func (p *ShellRunner) GetActions() []plugin.Action {
	logrus.Debug("Handling GetActions request!")
	return p.actions
}

func (p *ShellRunner) findActionByName(requestedName string) (*plugin.Action, error) {

	for _, action := range p.actions {
		if action.Name == requestedName {
			return &action, nil
		}
	}

	return nil, errors.New("unknown action was requested")
}

func (p *ShellRunner) executeActionEntryPoint(entryPointPath string, parameters []string) ([]byte, error) {

	logrus.Infoln("Executing entrypoint: ", entryPointPath, " with parameters: ", parameters)

	command := exec.Command(entryPointPath, parameters...)
	outputBytes, err := command.Output()
	if err != nil {
		logrus.Error("Failed to execute command with error: ", err)
		return nil, err
	}

	return outputBytes, nil
}

func (p *ShellRunner) ExecuteAction(request *plugin.ExecuteActionRequest) (*plugin.ExecuteActionResponse, error) {

	logrus.Debug("Handling ExecuteAction request")

	// Let's find the requested action.
	action, err := p.findActionByName(request.Name)
	if err != nil {
		logrus.Error("Invalid action was requested to be executed ", request.Name)
		return nil, err
	}

	// Parse the required parameters.
	var parameters []string
	for requestedParam, paramDescription := range action.Parameters {
		paramValue, ok := request.Parameters[requestedParam]
		if !ok && paramDescription.Required {
			err = errors.New("required parameter was not supplied: " + requestedParam)
			logrus.Error(err)
			return nil, err
		}

		parameters = append(parameters, paramValue)
	}

	// And finally execute the actual entrypoint.
	entryPointPath := path.Join(config2.GetConfig().Plugin.ActionsFolderPath, action.EntryPoint)
	ouputBytes, err := p.executeActionEntryPoint(entryPointPath, parameters)
	if err != nil {
		return nil, err
	}

	return &plugin.ExecuteActionResponse{ErrorCode: 0, Result: ouputBytes}, nil
}

func NewShellRunner(rootPluginDirectory string) (*ShellRunner, error) {
	config := config2.GetConfig()

	actions, err := actions2.LoadActionsFromDisk(path.Join(rootPluginDirectory, config.Plugin.ActionsFolderPath))
	if err != nil {
		return nil, err
	}

	description, err := description2.LoadPluginDescriptionFromDisk(path.Join(rootPluginDirectory, config.Plugin.PluginDescriptionFilePath))
	if err != nil {
		return nil, err
	}

	return &ShellRunner{
		actions:     actions,
		description: *description,
	}, nil
}