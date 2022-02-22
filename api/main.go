package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"willers-api/db"
	"willers-api/router"
)

// var hmacSecret = os.Getenv("SIGNINGKEY")

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	db.Init()
	defer db.Database.Close()

	e := router.Init()

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
