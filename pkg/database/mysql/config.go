package mysql

type MySqlConfig struct {
	DbName             string `yaml:"name"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	Database           string `yaml:"database"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	MaxIdleConnections int    `yaml:"maxIdleConnections"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
	Charset            string `yaml:"charset"`
	SSL                bool   `yaml:"ssl"`
	TimeZone           string `yaml:"time_zone"`
	Debug              bool   `yaml:"debug"`
}
