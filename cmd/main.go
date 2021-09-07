package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	budgetsaver "github.com/AndriyAntonenko/budgetSaver"
	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	"github.com/AndriyAntonenko/budgetSaver/pkg/handler"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
	service "github.com/AndriyAntonenko/budgetSaver/pkg/services"
	"github.com/spf13/viper"
)

func main() {
	cnf, err := config.InitAppConfig()
	if err != nil {
		log.Fatalf("error during config initialization: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	fileLogger := logger.InitFileLogger("Budget Saver", cnf.LogFile, cnf.Mode == "local")

	if err != nil {
		fileLogger.Error("postgresql initialization error", err, "func main()")
		panic(err)
	}

	srv := new(budgetsaver.Server)

	go func() {
		if err := srv.Run(cnf.Port, handlers.InitRoutes()); err != nil {
			fileLogger.Error("error during server running", err, "func main()")
			os.Exit(0)
		}
	}()

	fileLogger.Info("Server successfully started", "func main()")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// block main goroutine
	<-quit

	fileLogger.Info("Server shutting down", "func main()")

	if err := srv.Shutdown(context.Background()); err != nil {
		fileLogger.Error("error occured on server shutting down", err, "func main()")
		panic(err)
	}
}
