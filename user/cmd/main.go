package main

import (
	"user/config"
	"user/pkg/logger"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	var log logger.Interface
	l := logger.New(viper.GetString("server.domain"), 
					viper.GetString("log.level"))		
	log = l
	log.Debug("debug", "dfasdf")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
}