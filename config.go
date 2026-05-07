package main

import "os"

type Config struct {
	Address   string
	Token     string
	Path      string
	MountPath string
}

func ParseConfig() *Config {
	mountPath := os.Getenv("KV_MOUNT_PATH")
	if mountPath == "" {
		mountPath = "kv"
	}
	return &Config{
		Address:   os.Getenv("VAULT_ADDRESS"),
		Token:     os.Getenv("VAULT_TOKEN"),
		Path:      os.Getenv("SECRET_PATH"),
		MountPath: mountPath,
	}
}
