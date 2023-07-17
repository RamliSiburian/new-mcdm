package midleware

import (
	"context"
	"encoding/json"
	dto "mcdm/Dto"
	jwtToken "mcdm/Pkg/Mysql/Jwt"
	"net/http"
	"strings"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}
		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
