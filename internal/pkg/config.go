package pkg

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
)

const (
	path = "config/config.yaml"
)

type Telegram struct {
	Token string `yaml:"token"`
}

type OpenAI struct {
	Token string `yaml:"token"`
}

type AppConfig struct {
	Telegram
	OpenAI
}

var config *AppConfig

func ResolveConfig() (*AppConfig, error) {
	if config != nil {
		return config, nil
	}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content = []byte(os.ExpandEnv(string(content)))

	appConfig := AppConfig{}
	err = yaml.Unmarshal(content, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
