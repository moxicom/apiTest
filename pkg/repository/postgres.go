package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	SSLMode  string
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func InitDbConfig() Config {
	return Config{
		Host:     getEnv("host", "localhost"),
		User:     getEnv("user", "postgres"),
		Password: getEnv("password", ""),
		Dbname:   getEnv("dbname", "postgres"),
		Port:     getEnv("port", "5432"),
		SSLMode:  getEnv("sslmode", "disable"),
	}
}

func InitDb(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
		cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
