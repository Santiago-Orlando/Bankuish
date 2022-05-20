package services

import (
	"database/sql"
	"errors"
)

func CourseFinished(db *sql.DB, course string, firebaseID string) error {

	if course == "" {
		return errors.New("no course")
	}

	queryAddFinishedCourse := `INSERT INTO course_made_by_user ( user_id, course ) VALUES ( $1, $2 )`

	_, err := db.Query(queryAddFinishedCourse, firebaseID, course)
	if err != nil {
		return err
	}

	queryUpdateUserStudying := `UPDATE users SET studying = NULL WHERE firebase_id = $1`

	_, err = db.Query(queryUpdateUserStudying, firebaseID)
	if err != nil {
		return err
	}

	return nil
}
