package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type TsChromeExtTemplate struct {
}

var TsChromeExtTemplateInstance = NewTsChromeExtTemplate()

func NewTsChromeExtTemplate() *TsChromeExtTemplate {
	return &TsChromeExtTemplate{}
}

func (tlt *TsChromeExtTemplate) Url() string {
	return "https://github.com/pefish/create-chrome-ext-template.git"
}

func (tlt *TsChromeExtTemplate) Make(params global.TemplateParams) error {
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
		tlt.Url(),
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
