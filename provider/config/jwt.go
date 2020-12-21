package config

type JWT struct {
	SigningKey string `mapstructure:"signingKey" json:"signingKey" yaml:"signingKey"`
	Ttl        int64  `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
	RefreshTtl int64  `mapstructure:"refreshTtl" json:"refreshTtl" yaml:"refreshTtl"`
}
