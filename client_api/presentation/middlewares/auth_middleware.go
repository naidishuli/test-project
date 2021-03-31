package middlewares

import (
	"client_api/app/config"
	"client_api/app/errors"
	"client_api/domain/models"
	"client_api/presentation/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token, ok := verifyToken(authHeader)
		if !ok {
			response.ErrorResponse(w, r, errors.Unauthorized("invalid token"))
			return
		}

		user := token.Claims.(jwt.MapClaims)["user"].(models.User)
		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	})
}

func verifyToken(authValue string) (token *jwt.Token, ok bool) {
	bearerToken := strings.Split(authValue, " ")
	if len(bearerToken) != 2 {
		return
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error verification jwt method")
		}
		return []byte(config.Env.JwtVerificationKey), nil
	})

	if err != nil {
		return
	}

	if token.Valid {
		ok = true
	}

	return
}
