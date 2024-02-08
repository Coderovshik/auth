package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Coderovshik/auth/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.Get()

	path := fmt.Sprintf("file://%s", cfg.Migrations.MigrationsPath)
	url := fmt.Sprintf("%s/%s?x-migrations-collection=%s",
		cfg.MongoURI(),
		cfg.MongoDB.Database,
		cfg.Migrations.MigrationsCollection,
	)

	m, err := migrate.New(path, url)
	if err != nil {
		log.Fatalf("failed to create migrations object: %s", err.Error())
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	log.Println("migrations applied")
}
