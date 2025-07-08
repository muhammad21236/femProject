package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/muhammad21236/femProject/internal/app"
	"github.com/muhammad21236/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8086, "Port to run the application on")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	r := routes.SetupRoutes(application)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
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


