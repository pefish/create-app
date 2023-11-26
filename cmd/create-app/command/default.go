package command

import (
	"flag"
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	"github.com/pefish/create-app/pkg/templates"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
	go_prompt "github.com/pefish/go-prompt"
	"os"
	"os/exec"
	"strings"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	global.Templates = map[string]global.ITemplate{
		"electron":       templates.ElectronTemplateInstance,
		"go_app":         templates.GoAppTemplateInstance,
		"go_grpc_server": templates.GoGrpcServerTemplateInstance,
		"go_lib":         templates.GoLibTemplateInstance,
		"go_web_server":  templates.GoWebServerTemplateInstance,
		"react_app":      templates.ReactAppTemplateInstance,
		"react_dapp":     templates.ReactDappTemplateInstance,
		"rust_app":       templates.RustAppTemplateInstance,
		"rust_lib":       templates.RustLibTemplateInstance,
		"solidity_dapp":  templates.SolidityDappTemplateInstance,
		"ts_app":         templates.TsAppTemplateInstance,
		"ts_lib":         templates.TsLibTemplateInstance,
	}

	templateNames := make([]string, 0)
	for name, _ := range global.Templates {
		templateNames = append(templateNames, name)
	}
	flagSet.String("type", "", fmt.Sprintf("The template type. Available: [%s]", strings.Join(templateNames, ",")))
	flagSet.String("repo", "", "The repo url of project.")
	return nil
}

func (dc *DefaultCommand) Init(data *commander.StartData) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	if global.GlobalConfig.Type == "" {
		options := make([]go_prompt.InputOption, 0)
		for typeName, _ := range global.Templates {
			o := go_prompt.InputOption{
				Text: typeName,
				IsDefault: func() bool {
					return typeName == "go_app"
				}(),
			}

			options = append(options, o)
		}
		type_, isExit := go_prompt.PromptInstance.Input(
			"Please select type.",
			options,
		)
		if isExit {
			return nil
		}
		if type_ == "" {
			go_logger.Logger.InfoFRaw("error: required option '--type [string]' not specified.")
			return nil
		}
		global.GlobalConfig.Type = type_
	}
	if global.GlobalConfig.Repo == "" {
		repo, isExit := go_prompt.PromptInstance.Input(
			"Please select repo.",
			nil,
		)
		if isExit {
			return nil
		}

		if repo == "" {
			go_logger.Logger.InfoFRaw("error: required option '--repo [string]' not specified.")
			return nil
		}
		global.GlobalConfig.Repo = repo
	}

	atPos := strings.Index(global.GlobalConfig.Repo, "@")
	if atPos == -1 {
		go_logger.Logger.InfoFRaw("error: --repo [%s] is illegal.", global.GlobalConfig.Repo)
		return nil
	}
	colonPos := strings.Index(global.GlobalConfig.Repo, ":")
	slashPos := strings.Index(global.GlobalConfig.Repo, "/")
	domain := global.GlobalConfig.Repo[atPos+1 : colonPos]
	username := global.GlobalConfig.Repo[colonPos+1 : slashPos]
	projectName := global.GlobalConfig.Repo[slashPos+1 : len(global.GlobalConfig.Repo)-4]

	templateInstance, ok := global.Templates[global.GlobalConfig.Type]
	if !ok {
		go_logger.Logger.InfoFRaw("error: --type [%s] is illegal.", global.GlobalConfig.Type)
		return nil
	}
	params := global.TemplateParams{
		ProjectName: projectName,
		RepoUrl:     global.GlobalConfig.Repo,
		PackageName: fmt.Sprintf("%s/%s/%s", domain, username, projectName),
	}
	err := templateInstance.Make(params)
	if err != nil {
		return err
	}

	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
cd %s
git init
git remote add origin %s
`,
		params.ProjectName,
		params.RepoUrl,
	)
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
