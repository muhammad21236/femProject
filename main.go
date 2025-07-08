package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/muhammad21236/femProject/internal/app"
)

func main() {
	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	application.Logger.Println("Application running")

	http.HandleFunc("/health", healthCheck)
	server := &http.Server{
		Addr:         ":8086",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
