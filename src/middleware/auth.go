package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		secretKey := []byte(os.Getenv("AUTH_KEY"))

		if accessToken == "" {
			sendErrorResponse(w, "No authorization token provided", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(accessToken, "Bearer ") {
			sendErrorResponse(w, "Invalid authorization token format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(accessToken, "Bearer ")
		token, err := ParsingToken(secretKey, tokenString)
		if err != nil {
			var errMsg string
			if err.Error() == "Token is expired" {
				errMsg = "Token is expired"
			} else {
				errMsg = "Invalid token"
			}
			sendErrorResponse(w, errMsg, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			sendErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		subject, ok := claims["sub"].(string)
		if !ok {
			sendErrorResponse(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", subject)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

// ParsingToken melakukan parsing dan validasi token JWT
func ParsingToken(secretKey []byte, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
