package handlers

import (
	"Bankuish/api/config"
	"Bankuish/api/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func CoursesCompleted(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	firebaseID := r.Context().Value("firebaseID").(string)

	db := config.GetConnection()

	completedCourses, err := services.CoursesGetter(db, firebaseID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "err -> %s", err)
		return
	}

	data, err := json.Marshal(&completedCourses)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %s", err)
	}

	w.Write(data)
}
