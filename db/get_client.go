package db

import (
	"discord-metrics-server/v2/ent"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var entClient *ent.Client

func CreateClient() error {
	dataSourceString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))
	client, err := ent.Open("postgres", dataSourceString)
	entClient = client
	return err
}

func GetClient() *ent.Client {
	return entClient
}
