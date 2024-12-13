package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr        string
	PostgresAddress string
	MigrationsPath  string
}

func ReadConfig() Config {
	c := Config{}
	// Set default values
	c.HTTPAddr = "localhost:8080"
	c.PostgresAddress = "postgresql://postgres:password@localhost:5432/users?sslmode=disable"
	c.MigrationsPath = "file://internal/app/migrations"
	fmt.Println("reading config...")
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		c.HTTPAddr = httpAddr
	}
	dsn, exists := os.LookupEnv("DSN")
	if exists {
		c.PostgresAddress = dsn
	}
	migrationsPath, exists := os.LookupEnv("MIGRATIONS_PATH")
	if exists {
		c.MigrationsPath = migrationsPath
	}
	return c
}
