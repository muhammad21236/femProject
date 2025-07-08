package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/muhammad21236/femProject/internal/api"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "app: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Initialize the WorkoutHandler
	workoutHandler := api.NewWorkoutHandler()



	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}
