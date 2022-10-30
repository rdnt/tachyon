package main

import (
	"encoding/json"
	"os"
)

const path = "tachyon.json"

type Config struct {
	Redis Redis `json:"redis"`
}

type Redis struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Database  int    `json:"database"`
	StreamKey string `json:"streamKey"`
}

func (c Config) Save() error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		return createDefault()
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func createDefault() (Config, error) {
	c := Config{
		Redis: Redis{
			Host:      "0.0.0.0",
			Port:      6379,
			Database:  0,
			StreamKey: "events",
		},
	}

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return Config{}, err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
