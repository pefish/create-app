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
cat package.json | sed "s/template/%s/g" > temp && rm -rf package.json && mv temp package.json
npm install
`,
		rat.Url(),
		params.AppName,
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
