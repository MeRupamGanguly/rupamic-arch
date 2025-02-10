package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var seckey = "007JamesBond"

type contextKey string

const (
	UserIDKey contextKey = "userId"
	RolesKey  contextKey = "roles"
)

type Claims struct {
	UserId string   `json:"userId"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func CreateToken(id string, roles []string) (token string, err error) {
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{UserId: id, Roles: roles})
	token, err = tokenStr.SignedString([]byte(seckey))
	if err != nil {
		return "", err
	}
	return token, nil
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(seckey), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		if !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(*Claims)
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
