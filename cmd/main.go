package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	budgetsaver "github.com/AndriyAntonenko/budgetSaver"
	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/budgetSaver/pkg/router"
)

func initRouter() *router.Router {
	// Testing router
	r := router.NewRouter()

	r.Get("/api/:testParam/testStatic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test Get!!!")
	})

	r.Get("/api/testStatic/:testParam", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test Get 2!!!")
	})

	r.Get("/api/testStatic/1/:id/:token/sadasd/:pp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hard test")
	})

	return r
}

func main() {
	cnf, err := config.InitAppConfig()
	if err != nil {
		log.Fatalf("error during config initialization: %s", err.Error())
	}

	srv := new(budgetsaver.Server)

	r := initRouter()
	go func() {
		if err := srv.Run(cnf.Port, r); err != nil {
			log.Fatalf("error during server running: %s", err.Error())
		}
	}()

	fileLogger := logger.InitFileLogger("Budget Saver", cnf.LogFile, cnf.Mode == "local")
	fileLogger.Info("Server successfully started", "func main()")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// block main goroutine
	<-quit

	fileLogger.Info("Server shutting down", "func main()")
	fileLogger.Shutdown()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
