package config

type Oss struct {
	Local local `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}

type local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path" `
}

type qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath       string `mapstructure:"imgPath" json:"imgPath" yaml:"imgPath"`
	UseHTTPS      bool   `mapstructure:"useHttps" json:"useHttps" yaml:"useHttps"`
	AccessKey     string `mapstructure:"accessKey" json:"accessKey" yaml:"accessKey"`
	SecretKey     string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	UseCdnDomains bool   `mapstructure:"useCdnDomains" json:"useCdnDomains" yaml:"useCdnDomains"`
}
