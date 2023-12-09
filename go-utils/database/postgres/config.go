package postgres

import (
	"fmt"
	//_ "gopkg.in/yaml.v2"
)

type PostgresConfig struct {
	Host               string `mapstructure:"HOST"`
	Port               int    `mapstructure:"PORT"`
	Name               string `mapstructure:"NAME"`
	Username           string `mapstructure:"USERNAME"`
	Password           string `mapstructure:"PASSWORD"`
	MaxPoolSize        int    `mapstructure:"MAX_POOL_SIZE"`
	MaxIdleConnections int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
	SSLMode            string `mapstructure:"SSL_MODE"`
}

func GetPostgresConfig(port, poolSize, maxIdleConn int, user, password, host, dbName, sslMode string) PostgresConfig {
	return PostgresConfig{
		Host:               host,
		Port:               port,
		Name:               dbName,
		Username:           user,
		Password:           password,
		MaxPoolSize:        poolSize,
		MaxIdleConnections: maxIdleConn,
		SSLMode:            sslMode,
	}
}

func (cfg PostgresConfig) GetHost() string {
	return cfg.Host
}

func (cfg PostgresConfig) GetPort() int {
	return cfg.Port
}

func (cfg PostgresConfig) GetName() string {
	return cfg.Name
}

func (cfg PostgresConfig) GetMaxPoolSize() int {
	return cfg.MaxPoolSize
}

func (cfg PostgresConfig) GetMaxIdleConnections() int {
	return cfg.MaxIdleConnections
}

func (cfg PostgresConfig) GetSSLMode() string {
	return cfg.SSLMode
}

func (cfg PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s", cfg.Name, cfg.Username, cfg.Password, cfg.Host, cfg.SSLMode)
}

func (cfg PostgresConfig) GetConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
}
