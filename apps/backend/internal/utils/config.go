package utils

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	APIS       []API `yaml:"apis"`
	Server     `yaml:"server"`
	Components `yaml:"components"` // name of apis for each smd component
	Auth       struct {
		AccessSecKey  string `yaml:"access_sec_key"`
		RefreshSecKey string `yaml:"refresh_sec_key"`
	} `yaml:"auth"`
	Logger struct {
		LogLevel string `yaml:"log_level"`
	} `yaml:"logger"`
	Database struct {
		Name          string `yaml:"name"`
		Host          string `yaml:"host"`
		Port          int    `yaml:"port"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		Database      string `yaml:"database"`
		MigrationPath string `yaml:"migration_path"`
	} `yaml:"database"`
	Cors struct {
		AllowedOrigins []string `yaml:"allowed_origins"`
		AlloweMethods  []string `yaml:"allowed_methods"`
		AllowedHeaders []string `yaml:"allowed_headers"`
	} `yaml:"cors"`
	Prompts struct {
		TravelAI []struct {
			PromptType string `yaml:"prompt_type"`
			SystemRole string `yaml:"system_role"`
			UserPrompt string `yaml:"user_prompt"`
		} `yaml:"travel_ai"`
	} `yaml:"prompts"`
}

type API struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Key  string `yaml:"key"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Components struct {
	ItineraryGenerator string `yaml:"itinerary_generator"` // name of api like in a API struct
	AuthAPI            string `yaml:"auth_api"`            // name of api like in a API struct
}

func LoadConfig(confPath string) (*Config, error) {
	if confPath == "" {
		return nil, errors.New("config path is empty")
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(confPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
