package db

import (
	"Challenge/config"
	Log "Challenge/internal/adapters/Logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(config *config.Configurations) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Database.Dsn), &gorm.Config{})
	if err != nil {
		Log.Error.Printf("failed to connect database", err)
		return nil, err
	}
	Log.Info.Printf("Database is connected successfully ")
	return db, err
}

// GetDatabaseConnectionForTesting Connect to database for testing only
func GetDatabaseConnectionForTesting() (*gorm.DB, error) {
	var Configs = config.DatabaseConfigurations{
		Dsn: "postgresql://alisakr:2tcWfZK8kmhQnHUytUIFCA@free-tier4.aws-us-west-2.cockroachlabs.cloud:26257/tribaltransacion?sslmode=verify-full&options=--cluster%3Dslimy-hulk-3586",
	}
	db, err := gorm.Open(postgres.Open(Configs.Dsn), &gorm.Config{})
	if err != nil {
		Log.Error.Printf("5ff7624a-10e1-428c-a8bf-a413b1b745cd Failed to connect database in unit testing", err)
		return nil, err
	}
	Log.Info.Printf("Database is connected successfully in test mode in unit testing!")
	return db, err
}
