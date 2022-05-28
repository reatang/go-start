package demo

import (
	configModule "github.com/reatang/gostart/pkg/config"
	"github.com/reatang/gostart/pkg/database/mysql"
	"github.com/reatang/gostart/pkg/logger"
	"github.com/reatang/gostart/pkg/redis"
)

var config *Config

type Config struct {
	Web *configModule.Web `yaml:"web"`

	Database *mysql.MySqlConfig `yaml:"database"`
	Redis *redis.RedisConfig    `yaml:"redis"`

	Log *logger.LogConfig `yaml:"log"`
}

func InitConfig () {
	configModule.LoadConfig(&config)
}

func GetConfig() *Config {
	return config
}
