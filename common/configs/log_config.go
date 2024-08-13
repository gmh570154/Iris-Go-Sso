package configs

type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
	Save  uint   `yaml:"save"`
}
