package server

import (
	"api/config"
	"api/ent"
	"api/ent/migrate"
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func InitializeDB(cfg *config.Config) (db *ent.Client) {
	db = Connect(cfg)
	CreateDBSchema(db)
	return
}

func Connect(cfg *config.Config) *ent.Client {
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
		log.Errorf("Error connecting to the database: %v", err)
		os.Exit(1)
	}
	log.Info("Connected to the database.")

	return client
}

func CreateDBSchema(client *ent.Client) {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Errorf("Error creating schema resources: %v", err)
		os.Exit(1)
	}
	log.Info("Run auto-migration.")
}
