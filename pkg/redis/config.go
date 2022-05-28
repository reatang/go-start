package redis

type RedisConfig struct {
	Name     string   `yaml:"name"`
	Host     string   `yaml:"host"`
	Password string   `yaml:"password"`
	DB		 int      `yaml:"db"`
	MinIdle  int      `yaml:"minIdle"`
	PoolSize int      `yaml:"poolSize"`
}