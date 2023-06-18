package pkg

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	path = "config/config.yaml"
)

type Telegram struct {
	Token       string `yaml:"token"`
	Maintenance bool   `yaml:"maintenance"`
}

type OpenAI struct {
	Token string `yaml:"token"`
}

type AppConfig struct {
	Telegram
	OpenAI
}

func CreateConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
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
