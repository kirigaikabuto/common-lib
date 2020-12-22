package access_token_middleware

import (
	"context"
	"encoding/json"
	"net/http"
)

type AccessTokenMDW struct {
	accessTokenStore AccessTokenStore
}

type customError struct {
	Message string `json:"message"`
}

func NewAccessTokenMDW(accessTokenStore AccessTokenStore) (*AccessTokenMDW, error) {
	return &AccessTokenMDW{accessTokenStore}, nil
}

func (atmdw *AccessTokenMDW) Middleware(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		key := r.Header.Get("Authorization")
		if key == "" {
			respondJSON(w, http.StatusInternalServerError, &customError{"For access needed authorization header"})
			return
		}
		if key != "" {
			userId, err := atmdw.accessTokenStore.Get(key)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{"Not correct access key"})
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", userId)
			r = r.WithContext(ctx)
		}
		fn.ServeHTTP(w, r)
	})
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
