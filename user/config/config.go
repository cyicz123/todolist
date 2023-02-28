package config

import (
	"github.com/spf13/viper"
)

var (
	instance *viper.Viper
)

func init() {
	instance.SetConfigName("config")
	instance.SetConfigType("yml")
	instance.AddConfigPath("$HOME/.config/todolist/")
	err := instance.ReadInConfig()
	if err!=nil {
		panic(err)
	}
}

func GetInstance() *viper.Viper {
	return instance
}