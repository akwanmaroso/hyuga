package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()
	cfg := ParseConfig()

	client, err := vault.New(vault.WithAddress(cfg.Address), vault.WithRequestTimeout(30*time.Second))
	if err != nil {
		log.Fatal("failed to initialize client: ", err)
	}

	if err := client.SetToken(cfg.Token); err != nil {
		log.Fatal("failed to set token: ", err)
	}

	s, err := client.Secrets.KvV2Read(ctx, cfg.Path, vault.WithMountPath(cfg.MountPath))
	if err != nil {
		log.Fatal("failed to read secret: ", err)
	}

	f, err := os.Create(".env")
	if err != nil {
		log.Fatal("failed to create .env file: ", err)
	}
	defer f.Close()

	for key, value := range s.Data.Data {
		valueStr := fmt.Sprintf("%v", value)
		if s, ok := value.(string); ok {
			valueStr = fmt.Sprintf("%q", s)
		}
		if _, err := fmt.Fprintf(f, "%s=%s\n", key, valueStr); err != nil {
			log.Fatal("failed to write on .env: ", err)
		}
	}

	fmt.Println("Success generate .env")
}
