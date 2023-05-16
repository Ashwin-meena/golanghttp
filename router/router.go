package router

import (
	"go-dummy/controller"
	"go-dummy/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()
	publicRoute := router.PathPrefix("/public").Subrouter();
	publicRoute.HandleFunc("/api/newuser", controller.CreateUser).Methods("POST");
	
	privateRoute :=router.PathPrefix("/api").Subrouter();
    privateRoute.HandleFunc("/user/{id}", middleware.SetMiddlewareAuthentication(controller.GetUser)).Methods("GET")
    privateRoute.HandleFunc("/user", middleware.SetMiddlewareAuthentication(controller.GetAllUser)).Methods("GET")
    privateRoute.HandleFunc("/user/{id}",middleware.SetMiddlewareAuthentication( controller.UpdateUser)).Methods("PUT")
    publicRoute.HandleFunc("/api/login",controller.Login).Methods("GET")
    return router
}