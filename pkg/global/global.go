package global

type Config struct {
	Type string `json:"type"`
	Repo string `json:"repo"`
}

type TemplateParams struct {
	PackageName string `json:"package_name"`
	Username    string `json:"username"`
	AppName     string `json:"app_name"`
	RepoUrl     string `json:"repo_url"`
}

type ITemplate interface {
	Url() string
	Make(params TemplateParams) error
}

var Templates map[string]ITemplate

var GlobalConfig Config
