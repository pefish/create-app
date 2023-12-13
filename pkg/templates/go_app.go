package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
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
PACKAGE_NAME="%s" APP_NAME="%s" NAME="%s" ./init.sh
rm -rf ./init.sh
`,
		gat.Url(),
		params.AppName,
		params.AppName,
		params.PackageName,
		params.AppName,
		params.AppName,
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
