package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var instance *AppConfig
var once sync.Once

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

type AppConfig struct {
	Port     string
	LogFile  string
	Mode     string
	Postgres PostgresConfig
	Jwt      JwtConfig
}

func InitAppConfig() (*AppConfig, error) {
	var err error

	once.Do(func() {
		err = initConfig()
		if err != nil {
			return
		}

		err = godotenv.Load()
		if err != nil {
			return
		}

		instance = &AppConfig{
			Port:    viper.GetString("port"),
			LogFile: viper.GetString("logFile"),
			Mode:    os.Getenv("MODE"),
			Postgres: PostgresConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
				DBName:   os.Getenv("POSTGRES_DB"),
				SSLMode:  viper.GetString("db.sslMode"),
				Username: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
			},
			Jwt: JwtConfig{
				AccessTokenSecret:  os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
				RefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
			},
		}
	})

	return instance, err
}

func UseAppConfig() *AppConfig {
	if instance == nil {
		log.Fatalf("Config is not initialized")
	}

	return instance
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
