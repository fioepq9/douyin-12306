package config

type etcdConf struct {
	Address string `mapstructure:"address"`
}

func (etcdConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"Address": "0.0.0.0:2379",
	}
}
