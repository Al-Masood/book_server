package authMiddleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/al-masood/book_server/config"
	"github.com/golang-jwt/jwt/v5"
	"errors"
	"context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		if strings.HasPrefix(auth, "Basic") {
			userPassword := strings.TrimPrefix(auth, "Basic ")
			decodedUserPassword, err := base64.StdEncoding.DecodeString(userPassword)
			if err != nil {
				http.Error(w, "Error decoding username and password", http.StatusBadRequest)
				return
			}

			validUserPassword := config.APIConfig.AdminUser + ":" + config.APIConfig.AdminPassword
			if validUserPassword != string(decodedUserPassword) {
				http.Error(w, "Wrong username or password", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		} else if strings.HasPrefix(auth, "Bearer") {
			tokenString := strings.TrimPrefix(auth, "Bearer ")
			token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return config.APIConfig.ServerPrivateKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "jwtClaims", claims)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unsupported authentication method", http.StatusUnauthorized)
			return
		}
	})
}