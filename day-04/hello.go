package main

// 使用viper管理配置

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

func main() {
	fmt.Println("hello")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	var config Config
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Redis, config.MySQL)
}
