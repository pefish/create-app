package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
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
PACKAGE_NAME="%s" APP_NAME="%s" NAME="%s" ./init.sh
rm -rf ./init.sh
`,
		gwst.Url(),
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
