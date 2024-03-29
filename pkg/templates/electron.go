package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type ElectronTemplate struct {
}

var ElectronTemplateInstance = NewElectronTemplate()

func NewElectronTemplate() *ElectronTemplate {
	return &ElectronTemplate{}
}

func (et *ElectronTemplate) Url() string {
	return "https://github.com/pefish/create-electron-app-template.git"
}

func (et *ElectronTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
cd client && cat package.json | sed "s/template/%s/g" > temp.json && rm -rf package.json && mv temp.json package.json && cd ../
cd client/public
cat index.html | sed "s/template/%s/g" > temp.html && rm -rf index.html && mv temp.html index.html
cat package.json | sed "s/template/%s/g" > temp.json && rm -rf package.json && mv temp.json package.json
cd ../../
cd client && npm install
cp config/sample.yaml config/prod.yaml
cp config/sample.yaml config/local.yaml
`,
		et.Url(),
		params.AppName,
		params.AppName,
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
