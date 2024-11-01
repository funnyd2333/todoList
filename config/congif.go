package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	APP struct {
		Name string
		Port string
	}
	MYSQL struct {
		Dsn string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("unmarshal config err: %v", err)
	}
	InitDB()
}
