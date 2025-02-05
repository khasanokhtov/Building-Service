package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)


type Config struct{
	Database struct{
		Host		string `yaml:"host"`
		Port		string `yaml:"port"`
		User		string `yaml:"user"`
		Password	string `yaml:"password"`
		DB       	string `yaml:"db"`
	} `yaml:"database"`
}


var AppConfig Config

func LoadConfig(){
	file, err := os.Open("config/config.yaml")
	if err != nil{
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
	}
}