package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	configRelPath  = "../../"
	configFileName = "config"
	configFileType = "yml"
)

// C priority: ENV > config.yml > default
var C config

type config struct {
	Services servicesConf `mapstructure:"services"`
	Etcd     etcdConf     `mapstructure:"etcd"`
	Log      logConf      `mapstructure:"log"`
	MySQL    mysqlConf    `mapstructure:"mysql"`
	Redis    redisConf    `mapstructure:"redis"`
}

func init() {
	var err error

	// default config setting
	viper.SetDefault("services", servicesConf{}.defaultConf())
	viper.SetDefault("etcd", etcdConf{}.defaultConf())
	viper.SetDefault("log", logConf{}.defaultConf())
	viper.SetDefault("mysql", mysqlConf{}.defaultConf())
	viper.SetDefault("redis", redisConf{}.defaultConf())

	// config file setting
	viper.AddConfigPath(configRelPath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)

	// ENV config setting
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// read config
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// unmarshal to C
	err = viper.UnmarshalExact(&C)
	if err != nil {
		panic(err)
	}
}
