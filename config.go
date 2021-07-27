package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func (c PostgresConfig) Connection() string {
	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host: "localhost",
		Port: 5432,
		User: "piero",
		Name: "lenslocked_dev",
	}
}

type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Pepper   string         `json:"pepper"`
	HMACKey  string         `json:"hmac_key"`
	Database PostgresConfig `json:"database"`
}

func DefaultConfig() Config {
	return Config{
		Port:    3030,
		Env:     "dev",
		Pepper:  "secret-random-string",
		HMACKey: "secret-hmac-key",
	}
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func LoadConfig(configReq bool) Config {
	f, err := os.Open(".config")
	if err != nil {
		if configReq {
			log.Fatal(err)
		}
		log.Println("Using the default config")
		return DefaultConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Succesfully loaded .config")
	return c
}
