package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/muhammad21236/femProject/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8086, "Port to run the application on")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	

	http.HandleFunc("/health", healthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("Starting server on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
