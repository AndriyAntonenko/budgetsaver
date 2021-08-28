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

	"github.com/spf13/viper"
)

type BasicHandler struct{}

func (h *BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
}

func main() {

	// TODO: Implement my own logger
	if err := initConfig(); err != nil {
		log.Fatalf("error during config initialization: %s", err.Error())
	}

	srv := new(budgetsaver.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), &BasicHandler{}); err != nil {
			log.Fatalf("error during server running: %s", err.Error())
		}
	}()

	fmt.Println("Server successfully started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// block main goroutine
	<-quit

	fmt.Println("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
