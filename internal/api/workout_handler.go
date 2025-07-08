package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "workoutID")
	if paramsWorkoutID == "" {
		http.Error(w, "Workout ID is required", http.StatusBadRequest)
		return
	}
	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Fetching workout with ID: ", workoutID)

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	// Here you would typically parse the request body to get workout details
	// For simplicity, we will just return a success message
	fmt.Fprint(w, "Creating a new workout")
}