package handlers

import (
	"Bankuish/api/config"
	"Bankuish/api/models"
	"Bankuish/api/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func CourseInscription(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PATCH" {
		w.WriteHeader(405)
		return
	}

	firebaseID := r.Context().Value("firebaseID").(string)

	db := config.GetConnection()

	studying, err := services.StudyingGetter(db, firebaseID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "err -> %s", err)
		return
	}
	if studying != "" {
		w.WriteHeader(401)
		fmt.Fprintf(w, "You are currently in course: %s", studying)
		return
	}

	course := models.Course{}

	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "err -> %s", err)
		return
	}

	err = services.StudyingSetter(db, course.Course, firebaseID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "err -> %s", err)
		return
	}

	w.WriteHeader(202)
}
