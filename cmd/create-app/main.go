package main

import (
	"github.com/pefish/create-app/cmd/create-app/command"
	"github.com/pefish/create-app/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" is a tool to create various projects based on templates. author：pefish")
	commanderInstance.RegisterDefaultSubcommand("Use this command by default if you don't set subcommand.", command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
