package main

import "os"

type Config struct {
	Address string
	Token   string
	Path    string
}

func ParseConfig() *Config {
	return &Config{
		Address: os.Getenv("VAULT_ADDRESS"),
		Token:   os.Getenv("VAULT_TOKEN"),
		Path:    os.Getenv("SECRET_PATH"),
	}
}
