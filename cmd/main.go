package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const port = ":80"

func main() {
	log.Println("Starting service...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		for i := 15; i > 0; i-- {
			log.Println(i)
			time.Sleep(time.Second)
		}
	}))
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err)
			}
		}
	}()

	// waiting for interrupt signal
	<-ctx.Done()

	log.Println("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Bye bye...")
}
