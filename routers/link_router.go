package routers

import (
	"api.link/controllers"

	"github.com/gorilla/mux"
)

func InitItemRoutes(router *mux.Router) *mux.Router {

	subRouter := mux.NewRouter().PathPrefix("/link").Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.PostLink).Methods("POST")
	//subRouter.HandleFunc("{id}", controllers.PutItem).Methods("PUT")
	router.PathPrefix("/link").Handler(subRouter)

	router.HandleFunc("/{id}", controllers.GetLink).Methods("GET")
	return router
}
