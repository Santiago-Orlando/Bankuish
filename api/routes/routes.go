package routes

import (
	"net/http"

	"Bankuish/api/handlers"
	m "Bankuish/api/middlewares"
)

func Routes() {

	go http.HandleFunc("/register", handlers.Register)

	go http.HandleFunc("/courseOrder", m.TokenValidator(handlers.CourseOrderer))
	go http.HandleFunc("/courseInscription", m.TokenValidator(handlers.CourseInscription))
	go http.HandleFunc("/courseFinished", m.TokenValidator(handlers.CourseFinished))
	go http.HandleFunc("/courseCompleted", m.TokenValidator(handlers.CoursesCompleted))

}