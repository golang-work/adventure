package config

type Account struct {
	SubCountLimit int64 `mapstructure:"subCountLimit" json:"subCountLimit" yaml:"subCountLimit"`
	RecoverValidityHour uint `mapstructure:"recoverValidityHour" json:"recoverValidityHour" yaml:"recoverValidityHour"`
}
