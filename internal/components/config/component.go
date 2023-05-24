package config

import "github.com/spf13/viper"

type Config struct {
	config *viper.Viper
}

func New() *Config {
	return &Config{
		config: viper.New(),
	}
}

func (c *Config) SetDefaults() {
	c.config.SetDefault("app_port", 3000)
}

func (c *Config) AppPort() int {
	return c.config.GetInt("app_port")
}
