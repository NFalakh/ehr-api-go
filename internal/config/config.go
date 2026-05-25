package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ehr-api/internal/entity"
)

// Config holds all application configuration.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig reads configuration from environment variables.
func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "ehr_db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// DSN returns the MySQL Data Source Name.
func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

// InitDB initializes and returns a GORM database connection.
func (c *Config) InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto Migrate all entities
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Patient{},
		&entity.MedicalRecord{},
		&entity.Medication{},
		&entity.Allergy{},
		&entity.VitalSign{},
		&entity.LabResult{},
		&entity.Appointment{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
