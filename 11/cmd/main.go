package main

import (
	"11/internal/app"
	"11/internal/config"
	"11/internal/server"
	"11/internal/storage/memory"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	storage := memory.New()
	app := app.NewCalendar(storage)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := server.New(*cfg, app)
	go func() {
		log.Println("Server start...")
		if err := server.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server http stopped....")
			} else {
				log.Fatalln(err)
			}
		}
	}()
	<-stop
	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Shutdown(ctxClose)
	if err != nil {
		log.Fatalln(err)
	}
}
