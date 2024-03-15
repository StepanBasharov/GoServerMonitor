package types

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
)

type ServerAuth struct {
	Secret string
}

type ServerConfig struct {
	Host string
	Port int
	Auth ServerAuth
}

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
}

type SuperUserConfig struct {
	Username string
	Password string
	Email    string
}

type RedisConfig struct {
	Host string
	Port int
}

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Superuser SuperUserConfig
	Redis     RedisConfig
}

func (conf *Config) LoadConfig(configPath string) error {
	config, err := toml.LoadFile(configPath)
	if err != nil {
		log.Panicln(err)
		return err
	}

	err = config.Unmarshal(conf)

	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

func (serverConfig ServerConfig) GetBuildAddress() string {
	address := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	return address
}

func (databaseConfig DatabaseConfig) GetDsn() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		databaseConfig.Host,
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Database,
		databaseConfig.Port,
	)
	return dsn
}
