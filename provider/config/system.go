package config

type System struct {
	Debug    bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	DbType   string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	OssType  string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`
	Timezone string `mapstructure:"timezone" json:"timezone" yaml:"timezone"`
}
