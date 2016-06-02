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
	//port := "3000"

	//initialize routes
	router := routers.InitRoutes()
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	// initialize Database

	//local
	models.InitDB("postgres://swarup@localhost/LinkDB?sslmode=disable")
	models.InitRedis()
	defer models.RPool.Close()

	// heroku
	// models.InitDB(os.Getenv("DATABASE_URL"))
	// //redis
	// models.InitRedis()
	// defer models.RPool.Close()

	fmt.Print("starting server..")
	http.ListenAndServe(":"+port, n)
}
