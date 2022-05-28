package logger

type LogConfig struct {
	Path      string `yaml:"path"`
	Format    string `yaml:"format"`
	InfoFile  string `yaml:"infoFile"`
	ErrorFile string `yaml:"errorFile"`
}
