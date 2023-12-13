package templates

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/create-app/pkg/global"
)

type RustLibTemplate struct {
}

var RustLibTemplateInstance = NewRustLibTemplate()

func NewRustLibTemplate() *RustLibTemplate {
	return &RustLibTemplate{}
}

func (rlt *RustLibTemplate) Url() string {
	return "https://github.com/pefish/create-rust-lib-template.git"
}

func (rlt *RustLibTemplate) Make(params global.TemplateParams) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git clone %s --single-branch -v -b main --depth 1 %s
cd %s
rm -rf .git
cat Cargo.toml | sed "s/create-rust-lib-template/%s/g" > temp && rm -rf Cargo.toml && mv temp Cargo.toml
cargo update
`,
		rlt.Url(),
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
