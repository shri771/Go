package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUSerName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func (c *Config) SetUser(usreName string) error {
	c.CurrentUSerName = usreName
	return write(*c)

}

func Read() (Config, error) {

	filePath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	// Open and read the file
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	userConfig := Config{}
	err = json.NewDecoder(file).Decode(&userConfig)
	if err != nil {
		return Config{}, err
	}

	return userConfig, nil
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(home, configFileName)

	return filePath, nil
}

func write(c Config) error {

	filePath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(c)
	if err != nil {
		return err
	}
	return nil

}
