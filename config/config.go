package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GithubToken string
}

func New() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{
		GithubToken: viper.GetString("GITHUB_AUTH_TOKEN"),
	}

}
