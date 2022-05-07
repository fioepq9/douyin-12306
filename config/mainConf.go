package config

type mainConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (mainConf) defaultConf() map[string]string {
	return map[string]string{
		"host": "0.0.0.0",
		"port": "9090",
	}
}
