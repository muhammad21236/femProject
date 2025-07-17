package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/muhammad21236/femProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{workoutID}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workouts/{workoutID}", app.WorkoutHandler.HandleUpdateWorkoutByID)
	r.Delete("/workouts/{workoutID}", app.WorkoutHandler.HandleDeleteWorkoutByID)

	r.Post("/users", app.UserHandler.HandleRegisterUser)

	return r
}
