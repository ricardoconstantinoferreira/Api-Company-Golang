package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims["flag"].(string) == "user" {
			if !validUser(r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		if claims["flag"].(string) == "employee" {
			if !validEmployee(r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validEmployee(r *http.Request) bool {
	var url = getArrayUrls()
	var urlString string = r.RequestURI
	currentUrl := strings.Split(urlString, "/")
	iterator := 0

	for _, v := range url {
		vString := strings.Split(v, "/")
		if vString[1] == currentUrl[1] {
			iterator += 1
		}
	}

	if iterator > 0 {
		return false
	}

	return true
}

func validUser(r *http.Request) bool {

	var url = getArrayUrls()
	var urlString string = r.RequestURI
	currentUrl := strings.Split(urlString, "/")
	iterator := 0

	for _, v := range url {
		vString := strings.Split(v, "/")
		if vString[1] != currentUrl[1] {
			iterator += 1
		}
	}

	if iterator == len(url) {
		return false
	}

	return true
}

func getArrayUrls() [3]string {
	var url [3]string

	url[0] = "/create-products"
	url[1] = "/update-products-by-id"
	url[2] = "/delete-products-by-id"

	return url
}
