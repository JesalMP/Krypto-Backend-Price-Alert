package routes

import (
	"github.com/JesalMP/Krypto-Backend-Price-Alert/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/alerts/create/{alertPrice}", controllers.CreateAlert()).Methods("GET")
	router.HandleFunc("/alerts/delete/{alertPrice}", controllers.DeleteAlert()).Methods("DELETE")
	router.HandleFunc("/alerts", controllers.GetAllAlerts()).Methods("GET")
	router.HandleFunc("/alerts/{state}", controllers.GetAlertsState()).Methods("GET")
	//router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")

}
