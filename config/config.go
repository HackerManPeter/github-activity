package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	GithubToken string
}

func New() *Config {
	var githubToken string

	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		githubToken = os.Getenv("GITHUB_TOKEN")
	} else {
		viper.SetConfigFile(".env")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("fatal error config file: %v", err)
			os.Exit(1)
		}

		githubToken = viper.GetString("GITHUB_TOKEN")

	}

	if githubToken == "" {
		fmt.Printf("Set GITHUB_TOKEN as an environment variable")
		os.Exit(1)
	}

	return &Config{
		GithubToken: githubToken,
	}

}
