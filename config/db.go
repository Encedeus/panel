package config

import (
    "context"
    "entgo.io/ent/dialect/sql"
    "fmt"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/Encedeus/panel/ent/user"
    "github.com/Encedeus/panel/hashing"
    "github.com/labstack/gommon/log"
    _ "github.com/lib/pq"
)

func InitDB() *ent.Client {
    // Connect to database

    ctx := context.Background()

    db, err := ent.Open(
        "postgres",
        fmt.Sprintf(
            "host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
            Config.DB.Host,
            Config.DB.Port,
            Config.DB.User,
            Config.DB.DBName,
            Config.DB.Password,
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

    return db
}

func createSuperuser(db *ent.Client, ctx context.Context) {
    exists, err := db.User.Query().Where(user.Name("admin")).Exist(ctx)
    if err != nil {
        log.Fatalf("failed creating superuser: %e", err)
        return
    }
    if exists {
        return
    }

    userRole, err := db.Role.Query().Order(role.ByCreatedAt(sql.OrderAsc())).First(ctx)
    if err != nil {
        log.Fatalf("failed creating superuser: %e", err)
        return
    }

    _, err = db.User.Create().
        SetName("admin").
        SetPassword(hashing.HashPassword("admin")).
        SetEmail("admin@admin.com").
        SetRole(userRole).
        Save(ctx)

    if err != nil {
        log.Fatalf("failed creating superuser: %e", err)
        return
    }
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
