package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"

	"api.link/models"
	"api.link/routers"
)

func main() {

	//logger
	log.SetFormatter(&log.JSONFormatter{})

	//get the port
	port := os.Getenv("PORT")

	//initialize routes
	router := routers.InitRoutes()
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	// add static files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// initialize Database
	models.InitDB(os.Getenv("DATABASE_URL"))
	models.InitRedis()

	fmt.Print("starting server..")
	http.ListenAndServe(":"+port, n)
}
