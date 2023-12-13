package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type ReactAppTemplate struct {
}

var ReactAppTemplateInstance = NewReactAppTemplate()

func NewReactAppTemplate() *ReactAppTemplate {
	return &ReactAppTemplate{}
}

func (rat *ReactAppTemplate) Url() string {
	return "https://github.com/pefish/create-react-app-template.git"
}

func (rat *ReactAppTemplate) Make(params global.TemplateParams) error {
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
		rat.Url(),
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
