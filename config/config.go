package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func Load() *Config {
	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     "productos_db",
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
	log.Println("DBName cargado:", cfg.DBName)
	return cfg
}

func (c *Config) ConnectDB() (*sql.DB, error) {
	var dsn string
	if c.DBPassword == "" {
		dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
			c.DBUser, c.DBHost, c.DBPort, c.DBName)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	}

	log.Println("Conectando a la URL DSN:", dsn)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
