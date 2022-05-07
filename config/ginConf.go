package config

type ginConf struct {
	Mode string `mapstructure:"mode"`
}

func (ginConf) defaultConf() map[string]string {
	return map[string]string{
		"mode": "release",
	}
}
