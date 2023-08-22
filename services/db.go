package services

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/ent/role"
	"github.com/Encedeus/panel/ent/user"
	"github.com/Encedeus/panel/hashing"
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

	// db.Schema.Create(context.Background())

	// update Db schema
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// creates an admin user if it does not exist
	createSuperuserRole(db, ctx)
	createSuperuser(db, ctx)

	Db = db
	return db
}

func createSuperuser(db *ent.Client, ctx context.Context) {
	exists, err := db.User.Query().Where(user.Name("admin")).Exist(ctx)

	if exists {
		return
	}

	userRole, _ := db.Role.Get(ctx, 1)

	if err == nil && !exists {
		db.User.Create().
			SetName("admin").
			SetPassword(hashing.HashPassword("admin")).
			SetEmail("admin@admin.com").
			SetRole(userRole).
			Save(ctx)
		return
	}

	log.Error("failed to create superuser")
}
func createSuperuserRole(db *ent.Client, ctx context.Context) {
	exists, err := db.Role.Query().Where(role.Name("superuser")).Exist(ctx)

	if exists {
		return
	}

	if err == nil && !exists {
		_, err := db.Role.Create().
			SetName("superuser").
			SetPermissions([]string{"*"}).
			Save(ctx)
		log.Error(err)
		return
	}

	log.Error("failed to create superuser")
}
