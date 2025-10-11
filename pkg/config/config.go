package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yaml"

type Config struct {
	URLs []string `yaml:"users"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	data, err := os.ReadFile(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{
					URLs: []string{},
				},
				nil
		}
		return Config{}, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SaveConfigFile() error {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	finalFile := os.WriteFile(configFileName, data, 0644)

	return finalFile
}

func AddURL(url string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	cfg.URLs = append(cfg.URLs, url)

	return cfg.SaveConfigFile()
}

func RemoveURLByValue(index int) error {

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(cfg.URLs) {
		return fmt.Errorf("invalid index: %d", index)
	}

	cfg.URLs = append(cfg.URLs[:index], cfg.URLs[index+1:]...)

	return cfg.SaveConfigFile()
}
