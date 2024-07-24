package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type EnvConfigs struct {
	SECRET_KEY                 string
	ACCESS_TOKEN_EXPIRE_MINUES int
	DATABASE_HOST              string
	DATABASE_PORT              string
	DATABASE_USERNAME          string
	DATABASE_PASSWORD          string
	DATABASE_NAME              string
	DATABASE_SSL_MODE          string
}

var EnvConfig EnvConfigs

func InitConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file, %s", err))
	}

	viper.SetDefault("ACCESS_TOKEN_EXPIRE_MINUES", 15)

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&EnvConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
