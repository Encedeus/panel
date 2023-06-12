package service

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"panel/config"
	"panel/ent"
)

func DBInit() {
	fmt.Printf(
		"host=%s port=%d user=%s dbname=%s password=%s\n ",
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.User,
		config.Config.DB.DBName,
		config.Config.DB.Password)

	// Connect to database
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Config.DB.Host,
			config.Config.DB.Port,
			config.Config.DB.User,
			config.Config.DB.DBName,
			config.Config.DB.Password,
		),
	)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
