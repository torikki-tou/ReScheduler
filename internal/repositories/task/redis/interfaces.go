package redis

type Config interface {
	RedisHost() string
	RedisPort() int
}
