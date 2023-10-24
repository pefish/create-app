package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type GoLibTemplate struct {
}

var GoLibTemplateInstance = NewGoLibTemplate()

func NewGoLibTemplate() *GoLibTemplate {
	return &GoLibTemplate{}
}

func (glt *GoLibTemplate) Url() string {
	return "https://github.com/pefish/create-golang-lib-template.git"
}

func (glt *GoLibTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
NAME=%s ./init.sh
rm -rf ./init.sh
`,
		glt.Url(),
		params.ProjectName,
		params.ProjectName,
		params.PackageName,
	)
	err := go_shell.NewCmd(script).Run()
	if err != nil {
		return err
	}
	return nil
}
