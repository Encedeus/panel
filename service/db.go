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
		"host=%s port=%d user=%s dbname=%s password=%s\n ",
		config.Config.Db.Host,
		config.Config.Db.Port,
		config.Config.Db.User,
		config.Config.Db.DbName,
		config.Config.Db.Password)

	// connect to db

	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Config.Db.Host,
			config.Config.Db.Port,
			config.Config.Db.User,
			config.Config.Db.DbName,
			config.Config.Db.Password,
		),
	)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
