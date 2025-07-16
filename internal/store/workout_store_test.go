package store

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"testing"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5434 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}

	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("migrating test db: %v", err)
	}
	_, err = db.Exec(`TRUNCATE TABLE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("truncating test db: %v", err)
	}
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)
	tests := []struct {
		name     string
		workout  *Workout
		wantErr  bool
	}{
		{
			name: "Create valid workout",
			workout: &Workout{
				Title:              "Test Workout",
				Description: "This is a test workout",
				DurationMinutes:          60,
				CaloriesBurned:    200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Push-ups",
						Reps:         IntPtr(10),
						Sets:         3,
						Weight:       FloatPtr(135.5),
						Notes:        "Good form",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := store.CreteWorkout(tt.workout)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPtr(i float64) *float64 {
	return &i
}
