package config

import (
	"strconv"

	"github.com/Moosa-Razaa/Authentication/internal/helpers"
)

type DBConfig struct {
	Host                   string `mapstructure:"host"`
	Port                   int    `mapstructure:"port"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	DatabaseName           string `mapstructure:"database"`
	SSLMode                string `mapstructure:"sslmode"`
	MaxOpenConnections     int    `mapstructure:"max_open_connections"`
	MaxIdleConnections     int    `mapstructure:"max_idle_connections"`
	ConnectionsMaxLifetime int    `mapstructure:"connections_max_lifetime"`
}

func NewDatabaseConfig() *DBConfig {
	port, _ := strconv.Atoi(helpers.GetEnv("DB_PORT", "5432"))
	maxOpenConnections, _ := strconv.Atoi(helpers.GetEnv("DB_MAX_OPEN_CONNECTIONS", "10"))
	maxIdleConnections, _ := strconv.Atoi(helpers.GetEnv("DB_MAX_IDLE_CONNECTIONS", "5"))
	connectionsMaxLifetime, _ := strconv.Atoi(helpers.GetEnv("DB_CONNECTIONS_MAX_LIFETIME", "60"))

	return &DBConfig{
		Host:                   helpers.GetEnv("DB_HOST", "localhost"),
		Port:                   port,
		Username:               helpers.GetEnv("DB_USERNAME", "postgres"),
		Password:               helpers.GetEnv("DB_PASSWORD", "Password123_"),
		DatabaseName:           helpers.GetEnv("DB_NAME", "authdb"),
		SSLMode:                helpers.GetEnv("DB_SSL_MODE", "disable"),
		MaxOpenConnections:     maxOpenConnections,
		MaxIdleConnections:     maxIdleConnections,
		ConnectionsMaxLifetime: connectionsMaxLifetime,
	}
}

func (connectionString *DBConfig) GetConnectionString() string {
	return "host=" + connectionString.Host +
		" port=" + strconv.Itoa(connectionString.Port) +
		" user=" + connectionString.Username +
		" password=" + connectionString.Password +
		" dbname=" + connectionString.DatabaseName +
		" sslmode=" + connectionString.SSLMode
}

// docker run --name MoosaRaza -e POSTGRES_PASSWORD=Password123_ -p 5432:5432 -d postgres
