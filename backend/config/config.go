package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		panic(err)
	}

	InitDB()
}
