package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Bankuish/api/config"
	"Bankuish/api/models"

	"firebase.google.com/go/auth"
)


func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	client, err := config.Firebase()
	if err != nil {
		log.Fatalf("firebase error %v\n", err)
	}

	data := models.LoggingUser{}

	json.NewDecoder(r.Body).Decode(&data)

	user := (&auth.UserToCreate{}).
								Email(data.Email).
								Password(data.Password)

	newUser, err := client.CreateUser(context.Background(), user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "error %s", err)
		return
	}

	db := config.GetConnection()

	query := `INSERT INTO users ( firebase_id ) VALUES ( $1 )`

	_, err = db.Query(query, newUser.UID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error %s", err)
		return
	}

	w.WriteHeader(201)
}
