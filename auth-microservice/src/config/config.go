package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type UserMicroserviceConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type AuthMicroserviceConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ArticleMicroserviceConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type Config struct {
	DB      DBConfig                  `mapstructure:"db"`
	User    UserMicroserviceConfig    `mapstructure:"user-microservice"`
	Auth    AuthMicroserviceConfig    `mapstructure:"auth-microservice"`
	Article ArticleMicroserviceConfig `mapstructure:"article-microservice"`
	JWT     JWTConfig                 `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var configuration Config

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}

func LoadConfigByPath(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var configuration Config

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}
