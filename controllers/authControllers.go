package controllers

import (
	"api-go-deploy-heroku/models"
	u "api-go-deploy-heroku/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request
	if err != nil {
		u.Respond(w, u.Message("01", false, "Invalid request"))
		return
	}
	resp := account.Create() //create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode request body
	if err != nil {
		u.Respond(w, u.Message("01", false, "Invalid request"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
