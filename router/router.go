package router

import (
	"net/http"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/controller"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))

	userController := controller.UserController{}

	router.HandleFunc("/signup", userController.Signup()).Methods(http.MethodPost)
	router.HandleFunc("/login", userController.Login()).Methods(http.MethodPost)
	router.HandleFunc("/users", middleware.CheckAuth(userController.GetUsers())).Methods(http.MethodGet)
	router.HandleFunc("/users", middleware.CheckAuth(userController.Update())).Methods(http.MethodPut)

	return router
}
