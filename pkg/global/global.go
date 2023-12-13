package global

type Config struct {
	Type string `json:"type"`
	Repo string `json:"repo"`
}

type TemplateParams struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
	RepoUrl     string `json:"repo_url"`
}

type ITemplate interface {
	Url() string
	Make(params TemplateParams) error
}

var Templates map[string]ITemplate

var GlobalConfig Config
