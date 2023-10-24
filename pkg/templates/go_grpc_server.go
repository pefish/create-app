package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type GoGrpcServerTemplate struct {
}

var GoGrpcServerTemplateInstance = NewGoGrpcServerTemplate()

func NewGoGrpcServerTemplate() *GoGrpcServerTemplate {
	return &GoGrpcServerTemplate{}
}

func (ggst *GoGrpcServerTemplate) Url() string {
	return "https://github.com/pefish/create-golang-grpcserver-template.git"
}

func (ggst *GoGrpcServerTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
NAME="%s" ./init.sh
rm -rf ./init.sh
`,
		ggst.Url(),
		params.ProjectName,
		params.ProjectName,
		params.ProjectName,
	)
	err := go_shell.NewCmd(script).Run()
	if err != nil {
		return err
	}
	return nil
}
