package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/muhammad21236/femProject/internal/api"
	"github.com/muhammad21236/femProject/internal/store"
	"github.com/muhammad21236/femProject/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB // Assuming you want to keep a reference to the database connection
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "app: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Initialize the WorkoutHandler
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}
