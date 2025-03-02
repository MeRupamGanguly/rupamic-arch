package middlewares

import (
	"context"
	"net/http"
	"rupamic-arch/common"
)

func APPIKeyCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-KEY")
		if apiKey == "" {
			http.Error(w, common.ErrAPIKey.Error(), http.StatusBadRequest)
			return
		}

		role := common.GetAPIKeyFromDB(apiKey)
		if role == "" {
			http.Error(w, common.ErrAPIKey.Error(), http.StatusBadRequest)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "role", role))
		next.ServeHTTP(w, r)
	})
}
