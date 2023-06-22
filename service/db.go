package service

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"panel/config"
	"panel/ent"
	"panel/ent/user"
	"panel/util"
)

var Db *ent.Client

func InitDB() *ent.Client {
	// Connect to database

	ctx := context.Background()

	db, err := ent.Open(
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

	db.Schema.Create(context.Background())

	// update Db schema
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// creates an admin user if it does not exist
	createSuperuser(db, ctx)
	Db = db
	return db
}

func createSuperuser(db *ent.Client, ctx context.Context) {
	exists, err := db.User.Query().Where(user.Name("admin")).Exist(ctx)

	fmt.Println(exists)

	if exists {
		return
	}

	if err == nil && !exists {
		db.User.Create().
			SetName("admin").
			SetPassword(util.HashPassword("admin")).
			SetEmail("admin@admin.com").
			SetPfp([]byte{1, 2, 3}).
			Save(ctx)
		return
	}

	log.Error("failed to create superuser")
}
