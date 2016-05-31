package routers

import (
	"api.link/controllers"
	"api.link/middleware"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func InitAuthenticationRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/token", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.Handle("/refresh", negroni.New(
		negroni.HandlerFunc(middleware.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	router.HandleFunc("/something", controllers.Something).Methods("GET")
	return router
}
