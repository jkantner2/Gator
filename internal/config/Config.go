package config

import (
	"fmt"
	"os"
	"encoding/json"
	/*
	"bufio"
	*/
)

const configFileName = "/.gatorconfig.json"

//helper func obtain config filepath
func configFilepath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := homedir + configFileName
	return path, err
}

//read json and convert to struct
func ReadJSON() (Config, error){
	cfg := Config{}
	path, err := configFilepath()
	if err != nil {
		fmt.Printf("Error determining config path: %v\n", err)
		return cfg, err
	}
	
	content, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return cfg, err
		}

	defer content.Close()

	decoder := json.NewDecoder(content)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

//helper function to write to file

func write(cfg Config) error{
	path, err := configFilepath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

//set username and write config to json file
func (cfg *Config) SetUser (userName string) error{
	cfg.Current_username = userName
	return write(*cfg)
}
