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

func (pg *PostgresWorkoutStore) CreteWorkout(workout *Workout) (*Workout, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO workouts (title, description, duration_minutes, calories_burned)
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id`

	err = tx.QueryRow(query, workout.Title, workout.Description, workout.Duration, workout.CaloriesBurned).Scan(&workout.ID)
	if err != nil {
		return nil, err
	}
	for _, entry := range workout.Entries {
		entryQuery := `INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
					   VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					   RETURNING id`
		err = tx.QueryRow(entryQuery, workout.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&entry.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return workout, nil
}

func (pg *PostgresWorkoutStore) GetWorkoutByID(id int64) (*Workout, error) {
	Workout := &Workout{}
	return Workout, nil
}
