package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type GoWebServerTemplate struct {
}

var GoWebServerTemplateInstance = NewGoWebServerTemplate()

func NewGoWebServerTemplate() *GoWebServerTemplate {
	return &GoWebServerTemplate{}
}

func (gwst *GoWebServerTemplate) Url() string {
	return "https://github.com/pefish/create-golang-webserver-template.git"
}

func (gwst *GoWebServerTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
cat go.mod | sed "s/_template_/%s/g" > temp && rm -rf go.mod && mv temp go.mod
cp config/sample.yaml config/local.yaml
`,
		gwst.Url(),
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