package store

import "database/sql"

type Workout struct {
	ID             int            `json:"id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Duration       int            `json:"duration"`        // Duration in minutes
	CaloriesBurned int            `json:"calories_burned"` // Calories burned
	Entries        []WorkoutEntry `json:"entries"`         // Date of the workout in YYYY-MM-DD format
}

type WorkoutEntry struct {
	ID              int      `json:"id"`
	ExerciseName    string   `json:"exercise_name"`    // Name of the exercise
	Sets            int      `json:"sets"`             // Number of sets
	Reps            *int     `json:"reps"`             // Number of repetitions
	DurationSeconds *int     `json:"duration_seconds"` // Duration of the exercise in seconds
	Weight          *float64 `json:"weight"`           // Weight used in kg
	Notes           string   `json:"notes"`            // Additional notes for the entry
	OrderIndex      int      `json:"order_index"`      // Order index for sorting entries
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

type WorkoutStore interface {
	CreteWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
}
