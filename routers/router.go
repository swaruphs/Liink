package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = InitAuthenticationRoutes(router)
	router = InitItemRoutes(router)

	return router
}
