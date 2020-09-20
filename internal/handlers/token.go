package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/reedkihaddi/REST-API/pkg/models"
)

var jwtKey = []byte(os.Getenv("SECRET"))

// JWTAuthentication for authentication
func JWTAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// List of endpoints that  require auth
		auth := []string{"/products"}
		requestPath := r.URL.Path

		// Check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range auth {

			if value != requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			respondWithError(w, http.StatusForbidden, "Missing auth token")
			return
		}

		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			respondWithError(w, http.StatusForbidden, "Malformed authentication token")
			return
		}
		if !token.Valid {
			respondWithError(w, http.StatusForbidden, "Token is not valid.")
			return
		}

		next.ServeHTTP(w, r)

		/*
			IF WANT TO STORE JWT TOKEN IN HTTP COOKIE
			c, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					// If the cookie is not set, return an unauthorized status
					respondWithError(w, http.StatusUnauthorized, "Missing http cookie, go to /token")
					return
				}
				// For any other type of error, return a bad request status
				respondWithError(w, http.StatusBadRequest, "http cookie error")
				return
			}
			Get the JWT string from the cookie
			tknStr := c.Value
			claims := &models.Token{}
			tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					respondWithError(w, http.StatusUnauthorized, "signature invalid")
					return
				}
				respondWithError(w, http.StatusBadRequest, "error in jwt token")
				return
			}
			if !tkn.Valid {
				respondWithError(w, http.StatusUnauthorized, "token invalid")
				return
			}
			next.ServeHTTP(w, r) */
	})
}

// GetToken sets a JWT token
func GetToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expirationTime := time.Now().Add(1 * time.Minute)
		claims := &models.Token{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			respondWithError(w, http.StatusInternalServerError, "error in creating jwt token")
			return
		}
		/*
			IF WANT TO STORE IN HTTP COOKIE
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
				HttpOnly: true,
			})
			JWT Token */

		w.Write([]byte(tokenString))
	})
}
