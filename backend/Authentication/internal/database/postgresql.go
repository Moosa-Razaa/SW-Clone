package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Moosa-Razaa/Authentication/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func ExecuteMigrations() error {
	newDatabaseConfig := config.NewDatabaseConfig()
	connectionString := newDatabaseConfig.GetConnectionString()

	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	defer database.Close()

	err = database.Ping()
	if err != nil {
		return fmt.Errorf("error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Run Migrations
	if err := RunMigrations(database, newDatabaseConfig.DatabaseName); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully!")
	return nil
}

func RunMigrations(database *sql.DB, databaseName string) error {
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database migration driver: %w", err)
	}

	migrationsPath := "file://../../migrations"

	migrateInstance, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		return fmt.Errorf("failed to create new migration instance: %w", err)
	}

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	fmt.Println("Migrations applied successfully!")
	return nil
}

func InitializeDatabase() error {
	if db != nil {
		return nil // Already initialized
	}

	dbConfig := config.NewDatabaseConfig()

	var err error
	db, err = sql.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(dbConfig.MaxOpenConnections)
	db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnectionsMaxLifetime) * time.Second)

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	fmt.Println("Connected to the database successfully!")
	return nil
}

func GetDatabaseConnection() *sql.DB {
	if db == nil {
		if err := InitializeDatabase(); err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
	}
	return db
}

func CloseDatabase() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("failed to close the database: %w", err)
		}
		db = nil
		fmt.Println("Database connection closed successfully!")
	}
	return nil
}
