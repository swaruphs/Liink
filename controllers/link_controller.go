package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"api.link/helpers"
	"api.link/models"
	"api.link/services"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/mssola/user_agent"
	"github.com/speps/go-hashids"
)

func PostLink(w http.ResponseWriter, r *http.Request) {
	fmt.Print("inside post link")

	link := &models.Link{}

	err := r.ParseForm()
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(link, r.PostForm)

	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	id, err := services.PostLink(*link)
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	hd := hashids.NewData()
	hd.Salt = "linkdb-salt"
	hd.MinLength = 5
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{id})

	resp := map[string]string{}

	var schema string
	if r.TLS != nil {
		schema = "https://"
	} else {
		schema = "http://"
	}

	resp["url"] = schema + r.Host + "/" + e
	helpers.SendResponse(w, resp)

}

func GetLink(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sLinkID := vars["id"]

	u := user_agent.New(r.UserAgent())
	fmt.Print(u)

	log.WithFields(log.Fields{"id": sLinkID}).Warn("Get link with param")
	hd := hashids.NewData()
	hd.Salt = "linkdb-salt"
	hd.MinLength = 5
	h := hashids.NewWithData(hd)
	linkID, err := h.DecodeWithError(sLinkID)

	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	if linkID[0] == 0 {
		helpers.WriteBadRequestError(w, errors.New("linkID not found"))
		return
	}
	link, err := services.GetLink(linkID[0])
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	helpers.SendResponse(w, link)

	// if u.Platform() == "Android" {
	// 	http.Redirect(w, r, link.PlayStoreURL, 302)
	// 	return
	// }
	//
	// if u.Platform() == "iPhone" {
	// 	http.Redirect(w, r, link.AppstoreURL, 302)
	// 	return
	// }
	//
	// http.Redirect(w, r, link.WebURL, 302)

	//helpers.SendResponse(w, link)

}
