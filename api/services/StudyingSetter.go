package services

import (
	"database/sql"
)

func StudyingSetter(db *sql.DB, course string, firebaseID string) error {

	querySetStudying := `UPDATE users SET studying = $1 WHERE firebase_id = $2`

	_, err := db.Query(querySetStudying, course, firebaseID)
	if err != nil {
		return err
	}

	return nil
}
