package controller

import (
	"encoding/json"
	"jwt_auth/model"
	u "jwt_auth/util"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.WriteErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	err = account.Register() // Hesap yaratılır
	if err != nil {
		u.WriteErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	u.WriteJsonResponse(w, http.StatusCreated, account)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	if err != nil {
		u.WriteErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	err = account.Login() // Giriş yapılır
	if err != nil {
		u.WriteErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	u.WriteJsonResponse(w, http.StatusOK, account)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := model.GetAllAccounts()

	u.WriteJsonResponse(w, http.StatusOK, accounts)
}
