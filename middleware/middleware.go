package middleware

import (
	"fmt"
	"log"
	"net/http"

	"api.link/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

// Middleware function to check if the token is valid or not
func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return services.GetPublicKey(), nil
		}
	})

	if err == nil && token.Valid { //&& !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		value, ok := token.Claims["userid"]
		if ok == false {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		context.Set(req, "userid", int(value.(float64)))
		next(rw, req)
	} else {
		fmt.Println(err, token)
		rw.WriteHeader(http.StatusUnauthorized)
	}
}

//Recover handler - catches all panics and returns error message
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
