package utils

import (
	"confluence-payment/core-internal/conf"
	"reflect"

	// "strings"

	"github.com/spf13/viper"
)

var GlobalEnv conf.Env

func LoadGlobalEnv(path string) (err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	globalEnvType := reflect.TypeOf(conf.Env{})
	for i := 0; i < globalEnvType.NumField(); i++ {
		field := globalEnvType.Field(i)
		fieldTag := field.Tag.Get("mapstructure")
		if fieldTag != "" {
			viper.BindEnv(fieldTag)
		}
	}

	viper.Unmarshal(&GlobalEnv)
	return nil
}
