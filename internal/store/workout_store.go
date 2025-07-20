package store

import (
	"database/sql"
)

type Workout struct {
	ID              int            `json:"id"`
	UserID          int            `json:"user_id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"duration_minutes"` // Duration in minutes
	CaloriesBurned  int            `json:"calories_burned"`  // Calories burned
	Entries         []WorkoutEntry `json:"entries"`          // List of workout entries
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
	UpdateWorkout(*Workout) (*Workout, error)
	DeleteWorkout(id int64) error
	GetWorkoutOwner(id int64) (int, error)
}

func (pg *PostgresWorkoutStore) CreteWorkout(workout *Workout) (*Workout, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO workouts (user_id, title, description, duration_minutes, calories_burned)
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id`

	err = tx.QueryRow(query, workout.UserID, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned).Scan(&workout.ID)
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
	query := `SELECT id, title, description, duration_minutes, calories_burned
			  FROM workouts 
			  WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&Workout.ID, &Workout.Title, &Workout.Description, &Workout.DurationMinutes, &Workout.CaloriesBurned)
	if err == sql.ErrNoRows {
		return nil, nil // No workout found
	}
	if err != nil {
		return nil, err
	}

	entriesQuery := `SELECT id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index
					FROM workout_entries 
					WHERE workout_id = $1
					ORDER BY order_index`
	rows, err := pg.db.Query(entriesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		entry := WorkoutEntry{}
		err = rows.Scan(&entry.ID, &entry.ExerciseName, &entry.Sets, &entry.Reps, &entry.DurationSeconds, &entry.Weight, &entry.Notes, &entry.OrderIndex)
		if err != nil {
			return nil, err
		}
		Workout.Entries = append(Workout.Entries, entry)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Workout, nil
}

func (pg *PostgresWorkoutStore) UpdateWorkout(workout *Workout) (*Workout, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `UPDATE workouts 
			  SET title = $1, description = $2, duration_minutes = $3, calories_burned = $4 
			  WHERE id = $5`
	result, err := tx.Exec(query, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned, workout.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows // No workout found to update
	}
	_, err = tx.Exec(`DELETE FROM workout_entries WHERE workout_id = $1`, workout.ID)
	if err != nil {
		return nil, err
	}

	for _, entry := range workout.Entries {
		entryQuery := `INSERT INTO workout_entries 
		(workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err = tx.Exec(entryQuery, workout.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex)
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

func (pg *PostgresWorkoutStore) DeleteWorkout(id int64) error {
	query := `DELETE FROM workouts
	 WHERE id = $1`

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // No workout found to delete
	}

	return nil
}

func (pg *PostgresWorkoutStore) GetWorkoutOwner(workoutID int64) (int, error) {
	var userID int

	query := `SELECT user_id
	FROM workouts
	WHERE id=$1
	`

	err := pg.db.QueryRow(query, workoutID).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
