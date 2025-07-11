package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/muhammad21236/femProject/internal/store"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
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
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreteWorkout(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
	fmt.Println("Workout created successfully:", createdWorkout)
	w.WriteHeader(http.StatusCreated)
}
