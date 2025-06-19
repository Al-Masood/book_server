package handler

import (
	"github.com/al-masood/book_server/handler"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		auth := r.Header.Get("Authorization")

		if !strings.HasPrefix(auth, "Basic") && !strings.HasPrefix(auth, "Bearer"){
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		if strings.HasPrefix(auth, "Basic") {
			userpassword := strings.TrimPrefix(auth, "Basic ")
			
			decodedUserpassword, err := base64.StdEncoding.DecodeString(userpassword)

			if err != nil {
				http.Error(w, "Error decoding username and password", http.StatusBadRequest)
			}

			validUserpassword := handler.AdminUser + ":" + handler.AdminPassword

			if validUserpassword != string(decodedUserpassword) {
				http.Error(w, "Wrong username or password", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}

		if strings.HasPrefix(auth, "Bearer") {
			tokenStr := strings.TrimPrefix(auth, "Bearer ")

			_, err := handler.TokenAuth.Decode(tokenStr)
		
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusUnauthorized)
				return 
			}

			next.ServeHTTP(w, r)
		}

	})
}