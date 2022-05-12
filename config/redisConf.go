package config

type redisConf struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

func (redisConf) defaultConf() map[string]string {
	return map[string]string{
		"addr":     "0.0.0.0:6379",
		"username": "",
		"password": "root",
		"db":       "0",
	}
}
