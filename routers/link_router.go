package routers

import (
	"net/http"

	"api.link/controllers"

	"github.com/gorilla/mux"
)

func InitItemRoutes(router *mux.Router) *mux.Router {

	subRouter := mux.NewRouter().PathPrefix("/link").Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.PostLink).Methods("POST")
	//subRouter.HandleFunc("{id}", controllers.PutItem).Methods("PUT")
	router.PathPrefix("/link").Handler(subRouter)

	//add static files
	router.PathPrefix("/loaderio-b0fbe29c9838d25383e17790544f1a3f.txt").Handler(http.FileServer(http.Dir("./static/")))

	router.HandleFunc("/{id}", controllers.GetLink).Methods("GET")
	return router
}
