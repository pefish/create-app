package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
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
