package config

type Mysql struct {
	Driver       string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"dbName" json:"dbname" yaml:"dbName"`
	TablePrefix  string `mapstructure:"tablePrefix" json:"tablePrefix" yaml:"tablePrefix"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
	LogMode      bool   `mapstructure:"logMode" json:"logMode" yaml:"logMode"`
}
