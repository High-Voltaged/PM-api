package database

import (
	"api/config"
	"api/ent"
	"api/ent/migrate"
	"context"
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Initialize(cfg *config.Config) *ent.Client {
	// var DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", // for mysql
	var DSN = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	client, err := ent.Open("postgres", DSN)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Connected to the database.")

	return client
}

func CreateDBSchema(client *ent.Client) {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("Error creating schema resources: %v", err)
	}
	log.Println("Run auto-migration.")
}
