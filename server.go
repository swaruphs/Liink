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

	//initialize routes
	router := routers.InitRoutes()
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	// initialize Database
	models.InitDB(os.Getenv("DATABASE_URL"))
	models.InitRedis()

	fmt.Print("starting server..")
	http.ListenAndServe(":3000", n)
}
