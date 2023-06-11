package server

import (
	"Panel/ent"
	"context"
	"log"
)

func DBInit(db *ent.Client) {
	// Connect ot the database specified in the config file
	db, err := ent.Open("postgres", "")
	if err != nil {
		log.Fatalf("Failed connecting to the specified database: %v", err)
	}
	defer db.Close()
	// Run the auto migration tool to setup the database
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resourcers: %v", err)
	}
}
