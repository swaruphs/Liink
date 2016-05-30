package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"api.link/helpers"
	"api.link/models"
	"api.link/services"

	"github.com/gorilla/context"
	"github.com/gorilla/schema"
)

func Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := r.ParseForm()
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.PostForm)

	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	resp, err := services.Login(user.Name, user.Password)
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}
	helpers.SendResponse(w, resp)

}

func Something(w http.ResponseWriter, r *http.Request) {
	x := map[string]string{}
	x["hello"] = "world"
	helpers.SendResponse(w, x)
}

func Register(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	fmt.Println(r.PostForm)
	user := &models.User{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(user, r.PostForm)

	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	_, ierr := services.RegisterUser(user.Name, user.Password)
	if ierr != nil {
		helpers.WriteBadRequestError(w, ierr)
		return
	}

	helpers.SendEmptyResponse(w)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	userID, ok := context.GetOk(r, "userid")
	fmt.Println(userID)
	if ok == false {
		helpers.WriteBadRequestError(w, errors.New("unauthorized"))
		return
	}

	iuserID, ok := userID.(int)
	fmt.Println(iuserID)
	if ok == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp, err := services.RefreshToken(iuserID)
	if err != nil {
		helpers.WriteBadRequestError(w, err)
		return
	}

	helpers.SendResponse(w, resp)
}
