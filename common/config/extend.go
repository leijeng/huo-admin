package config

var Ext *Extend

type Extend struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type App struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
}
