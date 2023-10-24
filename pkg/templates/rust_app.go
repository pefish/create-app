package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type RustAppTemplate struct {
}

var RustAppTemplateInstance = NewRustAppTemplate()

func NewRustAppTemplate() *RustAppTemplate {
	return &RustAppTemplate{}
}

func (rat *RustAppTemplate) Url() string {
	return "https://github.com/pefish/create-rust-app-template.git"
}

func (rat *RustAppTemplate) Make(params global.TemplateParams) error {
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
		rat.Url(),
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
