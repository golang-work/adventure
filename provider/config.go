package provider

import "github.com/golang-work/adventure/provider/config"

type Config struct {
	Zap     config.Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   config.Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  config.System  `mapstructure:"system" json:"system" yaml:"system"`
	Account config.Account `mapstructure:"account" json:"account" yaml:"account"`
	Mysql   config.Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Oss     config.Oss     `mapstructure:"oss" json:"oss" yaml:"oss"`
	JWT     config.JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
