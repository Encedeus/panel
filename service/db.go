package service

import (
	"Panel/config"
	"Panel/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func DBInit() {
	fmt.Printf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.Config.Db.Host,
		config.Config.Db.Port,
		config.Config.Db.User,
		config.Config.Db.DbName,
		config.Config.Db.Password)

	// connect to db

	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s",
			config.Config.Db.Host,
			config.Config.Db.Port,
			config.Config.Db.User,
			config.Config.Db.DbName,
			config.Config.Db.Password,
		),
	)
	//client, err := ent.Open("postgres", "host=panel_db port=5432 user=Panel dbname=PanelDB password=testPassword")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
