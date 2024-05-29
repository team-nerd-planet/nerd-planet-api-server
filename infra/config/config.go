package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	App      App      `mapstructure:"APP"`
	Rest     Rest     `mapstructure:"REST"`
	Database Database `mapstructure:"DATABASE"`
	Swagger  Swagger  `mapstructure:"SWAGGER"`
	Jwt      Jwt      `mapstructure:"JWT"`
	Smtp     Smtp     `mapstructure:"SMTP"`
}

type App struct {
	Mode string `mapstructure:"MODE"` // "dev", "stage", "prod"
}

type Rest struct {
	Port int    `mapstructure:"PORT"`
	Mode string `mapstructure:"MODE"` // "debug" or "release"
}

type Database struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	LogLevel int    `mapstructure:"LOG_LEVEL"` // 1:Silent, 2:Error, 3:Warn, 4:Info
	UserName string `mapstructure:"USER_NAME"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"DB_NAME"`
}

type Swagger struct {
	Host     string `mapstructure:"HOST"`
	BasePath string `mapstructure:"BASE_PATH"`
}

type Jwt struct {
	SecretKey string `mapstructure:"SECRET_KEY"`
}

type Smtp struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	UserName string `mapstructure:"USER_NAME"`
	Password string `mapstructure:"PASSWORD"`
}

func NewConfig() (*Config, error) {
	_, b, _, _ := runtime.Caller(0)
	configDirPath := path.Join(path.Dir(b))

	conf := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDirPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Read config file.", "err", err)
		return nil, err
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println("Unmarshal config file.", "err", err)
		return nil, err
	}

	return &conf, nil
}
