package config

type mainConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (mainConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"host": "0.0.0.0",
		"port": "9090",
	}
}
