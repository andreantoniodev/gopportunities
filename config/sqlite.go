package config

import (
	"os"

	"github.com/andreantoniodev/gopportunities.git/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("SQLite database not found at %s, creating new database", dbPath)

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create SQLite database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate database schema: %v", err)
		return nil, err
	}

	return db, nil
}
