package config

import "time"

type mysqlConf struct {
	User   string       `mapstructure:"user"`
	Passwd string       `mapstructure:"passwd"`
	Host   string       `mapstructure:"host"`
	Port   string       `mapstructure:"port"`
	DBName string       `mapstructure:"db_name"`
	Log    mysqlLogConf `mapstructure:"log"`
}

func (mysqlConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"user":    "douyin_12306",
		"passwd":  "d1o2u3y0i6n",
		"host":    "0.0.0.0",
		"port":    "3306",
		"db_name": "douyin_12306",
		"log":     mysqlLogConf{}.defaultConf(),
	}
}

type mysqlLogConf struct {
	Level                     string        `mapstructure:"level"`
	Out                       string        `mapstructure:"out"`
	SlowThreshold             time.Duration `mapstructure:"slow_threshold"`
	IgnoreRecordNotFoundError bool          `mapstructure:"ignore_record_not_found_error"`
}

func (mysqlLogConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"level":                         "info",
		"out":                           "stdout",
		"slow_threshold":                500,
		"ignore_record_not_found_error": false,
	}
}
