package db

import (
	"discord-metrics-server/v2/ent"
	"fmt"
	"os"
)

func GetClient() (*ent.Client, error) {
	dataSourceString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))
	client, err := ent.Open("postgres", dataSourceString)
	return client, err
}
