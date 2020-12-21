package config

type Zap struct {
	Level    string `mapstructure:"level" json:"level" yaml:"level"`
	Format   string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director string `mapstructure:"director" json:"director"  yaml:"director"`
	ShowLine bool   `mapstructure:"showLine" json:"showLine" yaml:"showLine"`
}
