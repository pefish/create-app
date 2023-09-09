package templates

import (
	"fmt"
	"github.com/pefish/create-app/pkg/global"
	go_shell "github.com/pefish/go-shell"
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
		params.ProjectName,
		params.ProjectName,
		params.ProjectName,
	)
	err := go_shell.NewCmd(script).Run()
	if err != nil {
		return err
	}
	return nil
}
