package authMiddleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/al-masood/book_server/handler"

	"github.com/golang-jwt/jwt/v5"
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
			tokenString := strings.TrimPrefix(auth, "Bearer ")

		    token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return handler.ServerPrivateKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return 
			}

			next.ServeHTTP(w, r)
		}

	})
}