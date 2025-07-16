package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "Create valid workout",
			workout: &Workout{
				Title:           "Test Workout",
				Description:     "This is a test workout",
				DurationMinutes: 60,
				CaloriesBurned:  200,
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
		{
			name: "Create invalid workout",
			workout: &Workout{
				Title:           "Invalid Workout",
				Description:     "This is a test workout Invalid",
				DurationMinutes: 90,
				CaloriesBurned:  300,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Reps:         IntPtr(60),
						Sets:         3,
						Weight:       FloatPtr(135.5),
						Notes:        "form",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Reps:            IntPtr(12),
						Sets:            4,
						Weight:          FloatPtr(185),
						Notes:           "keep form",
						DurationSeconds: IntPtr(60),
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := store.CreteWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err, "Expected error for test case: %s", tt.name)
				return
			}
			require.NoError(t, err, "Unexpected error for test case: %s", tt.name)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err, "Failed to retrieve workout by ID for test case: %s", tt.name)

			assert.Equal(t, tt.workout.Title, retrieved.Title)
			assert.Equal(t, tt.workout.Description, retrieved.Description)
			assert.Equal(t, tt.workout.DurationMinutes, retrieved.DurationMinutes)
			assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

			for i := range retrieved.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Reps, retrieved.Entries[i].Reps)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].Weight, retrieved.Entries[i].Weight)
				assert.Equal(t, tt.workout.Entries[i].Notes, retrieved.Entries[i].Notes)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)
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
