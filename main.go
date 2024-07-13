package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
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

	s, err := client.Secrets.KvV2Read(ctx, cfg.Path, vault.WithMountPath("kv"))
	if err != nil {
		log.Fatal("failed to read secret: ", err)
	}

	f, err := os.Create(".env")
	if err != nil {
		log.Fatal("failed to create secret: ", err)
	}
	defer f.Close()

	for key, value := range s.Data.Data {
		valueStr := fmt.Sprintf("%v", value)
		if reflect.TypeOf(value).Kind() == reflect.String {
			valueStr = fmt.Sprintf("\"%v\"", value)
		}
		_, err := f.WriteString(fmt.Sprintf("%s=%s\n", key, valueStr))
		if err != nil {
			log.Fatal("failed to write on .env: ", err)
		}
	}

	fmt.Println("Success generate .env")
}
