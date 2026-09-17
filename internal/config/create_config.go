package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl    string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}
	var configStruct Config
	err = json.Unmarshal(content, &configStruct)
	if err != nil {
		return Config{}, err
	}
	return configStruct, nil
}

func (c *Config) SetUser(username string) error {
	c.UserName = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("can not write to file %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can not get home dir")
	}
	configFilePath := filepath.Join(homeDir, configFileName)
	return configFilePath, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	rawJsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, rawJsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
