package services

import "database/sql"

func StudyingGetter(db *sql.DB, firebaseID string) (string, error) {

	var studying sql.NullString

	queryGetStudying := `SELECT studying FROM users WHERE firebase_id = $1`

	err := db.QueryRow(queryGetStudying, firebaseID).Scan(&studying)
	if err != nil {
		return "", err
	}

	return studying.String, nil
}
