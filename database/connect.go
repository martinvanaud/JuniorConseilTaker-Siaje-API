package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Name     string `env:"POSTGRES_DB"`
}

func GetPostgresConnectionInfo(c PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

func GetPostgresConfig() PostgresConfig {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}

	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Port:     port,
		Name:     os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func Configure() {

	envParameters := GetPostgresConfig()
	config := GetPostgresConnectionInfo(envParameters)

	fmt.Println(config)

	db, postgresErr := gorm.Open(postgres.Open(config))

	if postgresErr != nil {
		log.Fatalf("Error could not initialize database: %v", postgresErr)
	} else {
		fmt.Println("Postgres database is ready")
	}

	DB = db
}
