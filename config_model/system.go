package config_model

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	HTTPS         bool   `mapstructure:"https" json:"https" yaml:"https"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
}
