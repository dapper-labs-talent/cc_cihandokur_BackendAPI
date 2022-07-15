package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/config"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/helper"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/middleware"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/model"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/service"
)

type UserController struct{}

var apiError = middleware.ApiError{}
var userService = service.UserService{}

func (u UserController) Signup() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var singup model.SingUp
		err := json.NewDecoder(r.Body).Decode(&singup)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := userService.CreateNewUser(&singup)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.ResponseWithJSON(w, resp)
	}
}

func (u UserController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var login model.Login
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := userService.Login(&login)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.ResponseWithJSON(w, resp)
	}
}

func (u UserController) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := userService.GetUsers()

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.ResponseWithJSON(w, resp)
	}
}

func (u UserController) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var update model.UserUpdate
		err := json.NewDecoder(r.Body).Decode(&update)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		headerToken := r.Header.Get(config.Config.Jwt.Header)

		claim, err := helper.ValidateToken(headerToken)
		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		err = userService.Update(claim.Email, &update)

		if err != nil {
			apiError.HandleErr(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.ResponseWithJSON(w, nil)
	}
}
