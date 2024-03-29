package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type TsAppTemplate struct {
}

var TsAppTemplateInstance = NewTsAppTemplate()

func NewTsAppTemplate() *TsAppTemplate {
	return &TsAppTemplate{}
}

func (tat *TsAppTemplate) Url() string {
	return "https://github.com/pefish/create-typescript-app-template.git"
}

func (tat *TsAppTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
PACKAGE_NAME="%s" APP_NAME="%s" USERNAME="%s" ./init.sh
rm -rf ./init.sh
`,
		tat.Url(),
		params.AppName,
		params.AppName,
		params.PackageName,
		params.AppName,
		params.Username,
	)
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
