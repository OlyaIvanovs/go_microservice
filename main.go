package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/OlyaIvanovs/go_microservice/handlers"
	"github.com/kelseyhightower/envconfig"
)

var opt struct {
	Port int `default:"9090"`
}

func main() {
	err := envconfig.Process("Microservice", &opt)
	if err != nil {
		log.Printf("Failed to parse command line arguments: %s", err.Error())
	}
	port := strconv.Itoa(opt.Port)
	log.Println(port)

	// Logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Handlers
	ph := handlers.NewProducts(l)

	// Create a new ServerMux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Create a new server
	s := &http.Server{
		Addr:         ":" + port,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
