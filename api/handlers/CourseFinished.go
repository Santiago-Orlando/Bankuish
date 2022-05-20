package handlers

import (
	"Bankuish/api/config"
	"Bankuish/api/services"
	"fmt"
	"net/http"
)

func CourseFinished(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PATCH" {
		w.WriteHeader(405)
		return
	}

	firebaseID := r.Context().Value("firebaseID").(string)

	db := config.GetConnection()

	course, err := services.StudyingGetter(db, firebaseID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error %s", err)
		return
	}
	if course == "" {
		w.WriteHeader(401)
		fmt.Fprintln(w, "You are not inscribed in any course!")
		return
	}

	err = services.CourseFinished(db, course, firebaseID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error %s", err)
		return
	}

}
