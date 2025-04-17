package config

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
