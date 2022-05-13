package config

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	configRelPath  = "./"
	configFileName = "config"
	configFileType = "yml"
)

// C priority: ENV > config.yml > default
var C config

type config struct {
	Main  mainConf  `mapstructure:"main"`
	Log   logConf   `mapstructure:"log"`
	Gin   ginConf   `mapstructure:"gin"`
	MySQL mysqlConf `mapstructure:"mysql"`
	Redis redisConf `mapstructure:"redis"`
}

func init() {
	var err error

	// default config setting
	viper.SetDefault("main", mainConf{}.defaultConf())
	viper.SetDefault("log", logConf{}.defaultConf())
	viper.SetDefault("gin", ginConf{}.defaultConf())
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
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}
