package config

import (
	"errors"
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

type ExchangeratesApiConfig struct {
	ApiKey string
	Url    string
}

type AppConfig struct {
	Port             string
	LogFile          string
	Mode             string
	Postgres         PostgresConfig
	ExchangeratesApi ExchangeratesApiConfig
	Jwt              JwtConfig
}

func InitAppConfig() (*AppConfig, error) {
	var err error

	once.Do(func() {
		err = initConfig()
		if err != nil {
			return
		}

		initEnv()

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
			ExchangeratesApi: ExchangeratesApiConfig{
				ApiKey: os.Getenv("EXCHANGE_SERVICE_API_KEY"),
				Url:    viper.GetString("externalServices.currencyExchange"),
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

func initEnv() {
	envFile := ".env.local"
	mode := os.Getenv("MODE")

	// if we runs locally
	if mode != "local" {
		// we will read env variables from os
		return
	}

	_, err := os.Stat(envFile)
	// if there no .env.local file just return
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	godotenv.Load(envFile)
}
