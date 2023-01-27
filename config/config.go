package config

import (
	"LogoForCardsGin/models"
	"encoding/json"
	"io"
	"log"
	"os"
)

func GetConfig() (*models.Config, error) {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var config models.Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &config, nil
}

// viper
