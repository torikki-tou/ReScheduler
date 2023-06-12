package config

import "github.com/spf13/viper"

const (
	StorageTypeRedis  = `redis`
	StorageTypeMemory = `memory`
)

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
	c.config.SetDefault("redis_port", 6379)
	c.config.SetDefault("redis_host", "redis")
	c.config.SetDefault("storage_type", StorageTypeRedis)
}

func (c *Config) AppPort() int {
	return c.config.GetInt("app_port")
}

func (c *Config) RedisPort() int {
	return c.config.GetInt("redis_port")
}

func (c *Config) RedisHost() string {
	return c.config.GetString("redis_host")
}

func (c *Config) GetStorageType() string {
	return c.config.GetString("storage_type")
}
