package services

import (
	"Bankuish/api/models"
	"database/sql"
)

func CoursesGetter(db *sql.DB, firebaseID string) ( models.CompletedCourses, error) {

	completedCourses := models.CompletedCourses{}

	queryGetCourses := `SELECT course FROM course_made_by_user WHERE user_id = $1`
	
	rows, err := db.Query(queryGetCourses, firebaseID)
	if err != nil {
		return completedCourses, err
	}

	for rows.Next() {
		var course models.Course

		err = rows.Scan(&course.Course)
		if err != nil {
			return completedCourses, err
		}

		completedCourses.Courses = append(completedCourses.Courses, course)

	}

	return completedCourses, nil
}
