package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type TsLibTemplate struct {
}

var TsLibTemplateInstance = NewTsLibTemplate()

func NewTsLibTemplate() *TsLibTemplate {
	return &TsLibTemplate{}
}

func (tlt *TsLibTemplate) Url() string {
	return "https://github.com/pefish/create-typescript-lib-template.git"
}

func (tlt *TsLibTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
cat package.json | sed "s/template_name/%s/g" > temp.json && rm -rf package.json && mv temp.json package.json
npm install
`,
		tlt.Url(),
		params.ProjectName,
		params.ProjectName,
		params.ProjectName,
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
