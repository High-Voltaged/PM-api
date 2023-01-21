package database

import (
	"api/config"
	"api/ent"
	"api/ent/migrate"
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Initialize(cfg *config.Config) *ent.Client {
	var DSN = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Name,
		cfg.Database.Password,
	)

	client, err := ent.Open("postgres", DSN)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Connected to the database.")

	return client
}

func Migrate(client *ent.Client) {
	ctx := context.Background()

	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("Error creating schema resources: %v", err)
	}
	log.Println("Run auto-migration.")
}
