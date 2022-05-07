package config

type logConf struct {
	Level string `mapstructure:"level"`
	Out   string `mapstructure:"out"`
}

func (logConf) defaultConf() map[string]string {
	return map[string]string{
		"level": "warn",
		"out":   "stdout",
	}
}
