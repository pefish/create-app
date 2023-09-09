package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
)

type SolidityDappTemplate struct {
}

var SolidityDappTemplateInstance = NewSolidityDappTemplate()

func NewSolidityDappTemplate() *SolidityDappTemplate {
	return &SolidityDappTemplate{}
}

func (sdt *SolidityDappTemplate) Url() string {
	return "https://github.com/pefish/create-solidity-dapp-template.git"
}

func (sdt *SolidityDappTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
npm install
`,
		sdt.Url(),
		params.ProjectName,
		params.ProjectName,
	)
	err := go_shell.NewCmd(script).Run()
	if err != nil {
		return err
	}
	return nil
}
