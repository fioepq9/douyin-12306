package config

type ginConf struct {
	Mode    string `mapstructure:"mode"`
	Timeout int64  `mapstructure:"timeout"`
}

func (ginConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"mode":    "release",
		"timeout": 5,
	}
}
