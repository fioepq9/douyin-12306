package config

type mysqlConf struct {
	User   string `mapstructure:"user"`
	Passwd string `mapstructure:"passwd"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DBName string `mapstructure:"db_name"`
}

func (mysqlConf) defaultConf() map[string]string {
	return map[string]string{
		"user":    "douyin_12306",
		"passwd":  "d1o2u3y0i6n",
		"host":    "0.0.0.0",
		"port":    "3306",
		"db_name": "douyin_12306",
	}
}
