package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"history-guesser/server/internal/config"
	"history-guesser/server/internal/db"
	"history-guesser/server/internal/rounds"
	"history-guesser/server/internal/websocket"
)

func main() {
	cfg := config.Load()
	database, err := db.New(cfg.DatabaseURL)
	if err != nil { log.Fatal(err) }
	defer database.Close()

	roundRepo := rounds.NewPGRepository(database)
	hub := websocket.NewHub(roundRepo)
	r := websocket.Router(hub)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: r}
	go func() {
		log.Printf("api listening on %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
