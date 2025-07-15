package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	js = append(js, '\n') // Add newline for readability
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ReadIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "workoutID")
	if idParam == "" {
		return 0, errors.New("missing id parameter")
	}
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
