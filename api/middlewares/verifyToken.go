package middlewares

import (
	"Bankuish/api/config"
	"context"
	"fmt"
	"net/http"
)

func TokenValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client, err := config.Firebase()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error %s", err)
			return
		}

		token := r.Header.Get("Authorization")

		authToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			w.WriteHeader(401)
			fmt.Fprintf(w, "error %s", err)
			return
		}

		ctxValues := context.WithValue(r.Context(), "firebaseID", authToken.UID)
		r = r.WithContext(ctxValues)
		
		next(w, r)
	})
}
