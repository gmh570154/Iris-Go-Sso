package models

type Config struct {
	Sso      *Sso     `yaml:"Sso"`
	Console  *Console `yaml:"Console"`
	LogLevel string   `yaml:"Log_level"`
	Videos   []Video  `yaml:"Video"`
}

type Sso struct {
	LoginUrl  string `yaml:"login_url"`
	LogoutUrl string `yaml:"logout_url"`
	HostUrl   string `yaml:"host_url"`
	GrantType string `yaml:"grant_type"`
	ClientId  string `yaml:"client_id"`
	SecretKey string `yaml:"secret_key"`
}

type Console struct {
	HomeUrl string `yaml:"home_url"`
}

type Video struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type UserVar struct {
	LoginUrl  string `yaml:"login_url"`
	LogoutUrl string `yaml:"logout_url"`
	GrantType string `yaml:"grant_type"`
	HomeUrl   string `yaml:"home_url"`
}
