package config

import "time"

type servicesConf struct {
	Api  apiServiceConf  `mapstructure:"api"`
	User userServiceConf `mapstructure:"user"`
}

func (servicesConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"api":  apiServiceConf{}.defaultConf(),
		"user": userServiceConf{}.defaultConf(),
	}
}

type apiServiceConf struct {
	Name    string        `mapstructure:"name"`
	Addr    string        `mapstructure:"addr"`
	Mode    string        `mapstructure:"mode"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func (apiServiceConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"name":    "api",
		"addr":    "0.0.0.0:9090",
		"mode":    "release",
		"timeout": 5,
	}
}

type userServiceConf struct {
	Name   string                `mapstructure:"name"`
	Addr   string                `mapstructure:"addr"`
	Client userServiceClientConf `mapstructure:"Client"`
	Server userServiceServerConf `mapstructure:"Server"`
}

func (userServiceConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"name":   "user",
		"addr":   "0.0.0.0:9091",
		"client": userServiceClientConf{}.defaultConf(),
		"server": userServiceServerConf{}.defaultConf(),
	}
}

type userServiceClientConf struct {
	MuxConnection int           `mapstructure:"mux_connection"`
	RpcTimeout    time.Duration `mapstructure:"rpc_timeout"`
	ConnTimeout   time.Duration `mapstructure:"conn_timeout"`
}

func (userServiceClientConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"mux_connection": 1,
		"rpc_timeout":    3,
		"conn_timeout":   50,
	}
}

type userServiceServerConf struct {
	MaxConnections int `mapstructure:"max_connections"`
	MaxQPS         int `mapstructure:"max_qps"`
}

func (userServiceServerConf) defaultConf() map[string]interface{} {
	return map[string]interface{}{
		"max_connections": 1000,
		"max_qps":         100,
	}
}
