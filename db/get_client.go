package db

import (
	"discord-metrics-server/v2/ent"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var entClient *ent.Client

func CreateClient() error {
	dataSourceString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("MARIA_USER"),
		os.Getenv("MARIA_PASSWORD"),
		os.Getenv("MARIA_HOST"),
		os.Getenv("MARIA_PORT"),
		os.Getenv("MARIA_DB"))
	fmt.Println(dataSourceString)
	client, err := ent.Open("mysql", dataSourceString)
	if err != nil {
		log.Fatalf("Error opening mariadb connection!")
	}
	entClient = client
	return err
}

func GetClient() *ent.Client {
	return entClient
}
