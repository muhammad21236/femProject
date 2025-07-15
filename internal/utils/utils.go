package utils

import (
	"encoding/json"
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
	js = append(js, '\n') // Add a newline for better readability
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil

}

func ReadIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		return 0, http.ErrNoLocation
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}
