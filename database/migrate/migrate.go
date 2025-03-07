package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	migrationType := flag.String("t", "", "Migration Type (Value: 'up' & 'down')")
	flag.Parse()
	if *migrationType == "" {
		log.Fatal("Migration Type (-t) is required flag. Value: 'up' & 'down'.")
	}
	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = sqlDB.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrateWithDatabaseInstance, err := migrate.NewWithDatabaseInstance("file://database/migrate", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	switch *migrationType {
	case "down":
		if err = migrateWithDatabaseInstance.Steps(-1); err != nil && err.Error() != "no change" {
			log.Fatal(err)
		}
	case "up":
		if err = migrateWithDatabaseInstance.Up(); err != nil && err.Error() != "no change" {
			log.Fatal(err)
		}
	}
}
