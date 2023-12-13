package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
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
PACKAGE_NAME="%s" APP_NAME="%s" NAME="%s" ./init.sh
rm -rf ./init.sh
`,
		glt.Url(),
		params.AppName,
		params.AppName,
		params.PackageName,
		params.AppName,
		params.PackageName,
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
