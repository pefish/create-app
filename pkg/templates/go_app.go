package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type GoAppTemplate struct {
}

var GoAppTemplateInstance = NewGoAppTemplate()

func NewGoAppTemplate() *GoAppTemplate {
	return &GoAppTemplate{}
}

func (gat *GoAppTemplate) Url() string {
	return "https://github.com/pefish/create-golang-app-template.git"
}

func (gat *GoAppTemplate) Make(params global.TemplateParams) error {
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
		gat.Url(),
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
