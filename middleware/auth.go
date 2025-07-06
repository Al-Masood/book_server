package authMiddleware

import (
	"github.com/al-masood/book_server/domain/service"
	"net/http"
	"strings"
)

func AuthMiddleware(userService service.UserService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		if strings.HasPrefix(auth, "Basic") {
			credentials := strings.TrimPrefix(auth, "Basic ")

			err := userService.AuthenticateBasic(ctx, credentials)

			if err != nil {
				http.Error(w, "Authentication failed", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		} else if strings.HasPrefix(auth, "Bearer") {
			token := strings.TrimPrefix(auth, "Bearer ")

			err := userService.AuthenticateBearer(ctx, token)

			if err != nil {
				http.Error(w, "Authentication failed", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unsupported authentication method", http.StatusUnauthorized)
			return
		}
	})
}
