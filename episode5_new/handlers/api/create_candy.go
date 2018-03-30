package api

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CreateCandy(db models.DB) http.Handler {
	type jsonInput struct {
		Name string `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
