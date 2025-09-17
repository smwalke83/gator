package config

import (
	"encoding/json"
	"os"
	"fmt"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL				string `json:"db_url"`
	CurrentUserName		string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error{
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	var c Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c, err
	}
	filePath := filepath.Join(homeDir, configFileName)
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c, err
	}
	err = json.Unmarshal(fileBytes, &c)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		return c, err
	}
	return c, nil
}

func write(c Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	filePath := filepath.Join(homeDir, configFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	return nil
}