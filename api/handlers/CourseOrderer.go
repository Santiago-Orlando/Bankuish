package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Bankuish/api/models"
	"Bankuish/api/services"
)

func CourseOrderer(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		w.WriteHeader(405)
		return
	}

	courses := models.Courses{}

	err := json.NewDecoder(r.Body).Decode(&courses)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error %s", err)
	}

	orderedCourses := services.CourseOrdererService(courses)

	data, err := json.Marshal(&orderedCourses)
	if err != nil {
		w.WriteHeader(401)
		fmt.Fprintln(w, "error: empty array")
	}

	w.Write(data)
}
