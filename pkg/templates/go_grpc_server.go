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
cat go.mod | sed "s/_template_/%s/g" > temp && rm -rf go.mod && mv temp go.mod
cat main.go | sed "s/_template_/%s/g" > temp && rm -rf main.go && mv temp main.go
cat client/main.go | sed "s/_template_/%s/g" > client/temp && rm -rf client/main.go && mv client/temp client/main.go
cat service/helloworld/helloworld.go | sed "s/_template_/%s/g" > service/helloworld/temp && rm -rf service/helloworld/helloworld.go && mv service/helloworld/temp service/helloworld/helloworld.go
cp config/sample.yaml config/local.yaml
`,
		ggst.Url(),
		params.ProjectName,
		params.ProjectName,
		params.ProjectName,
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
