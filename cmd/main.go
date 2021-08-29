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

	r.Post("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method : %s Serving: %s\n", "POST", r.URL.Path)
	})

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Working!!!")
	})

	return r
}

func main() {
	cnf, err := config.InitAppConfig()
	if err != nil {
		log.Fatalf("error during config initialization: %s", err.Error())
	}

	srv := new(budgetsaver.Server)

	go func() {
		if err := srv.Run(cnf.Port, initRouter()); err != nil {
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
