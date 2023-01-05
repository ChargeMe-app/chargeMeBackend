package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var cfg Config

type Config struct {
	Postgres      Database      `yaml:"DB"`
	YandexStorage YandexStorage `yaml:"YandexStorage"`
	Sitronics     Sitronics     `yaml:"Sitronics"`
	MyECars       MyECars       `yaml:"MyECars"`
}

type Sitronics struct {
	BaseUrl string `yaml:"baseUrl"`
}

type MyECars struct {
	BaseUrl string `yaml:"baseUrl"`
	Key     string `json:"key"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

type YandexStorage struct {
	BaseUrl string `yaml:"baseUrl"`
}

func InitConfig() (*Config, error) {
	configPath := "cmd/chargeMe/config.yaml"

	clean := filepath.Clean(configPath)

	file, err := os.Open(clean)
	if err != nil {
		return nil, fmt.Errorf("fail to open config file in path \"%s\" with error %w", configPath, err)
	}

	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to parse config %w", err)
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
