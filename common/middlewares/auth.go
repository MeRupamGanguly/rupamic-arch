package middlewares

import (
	"context"
	"net/http"
	"rupamic-arch/common/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "userId"
	RolesKey  contextKey = "roles"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenStr, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.Seckey), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		if !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(*auth.Claims)
		if !ok {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, RolesKey, claims.Roles)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
