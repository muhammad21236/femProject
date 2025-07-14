package api

import (
	"encoding/json"
	"fmt"
	"log"
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

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch the workout", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workout)

	// fmt.Fprint(w, "Fetching workout with ID: ", workoutID)

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		log.Println("failed to create workout:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request) {
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

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		http.Error(w, "Failed to fetch workout", http.StatusInternalServerError)
		return
	}
	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"` // Duration in minutes
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}
	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}
	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}
	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}
	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}
	updatedWorkout, err := wh.workoutStore.UpdateWorkout(existingWorkout)
	if err != nil {
		fmt.Println("update workout error", err)
		http.Error(w, "Failed to update workout", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedWorkout)

}
