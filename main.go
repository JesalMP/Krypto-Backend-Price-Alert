package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/JesalMP/Krypto-Backend-Price-Alert/configs"
	"github.com/JesalMP/Krypto-Backend-Price-Alert/controllers"
	"github.com/JesalMP/Krypto-Backend-Price-Alert/routes"
	"github.com/gorilla/mux"
)

func theRouter(wg *sync.WaitGroup) {
	println("Therouter")
	defer wg.Done()
	router := mux.NewRouter()
	//configs.ConnectDB()
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":6061", router))
	wg.Done()
}
func Mailer(wg *sync.WaitGroup) {

	defer wg.Done()
	for {
		println("waiting for trigger")
		controllers.Trigger()
	}
	wg.Done()
	//controllers.Trigger()

}
func main() {

	//run database
	configs.ConnectDB()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go theRouter(wg)
	go Mailer(wg)
	wg.Wait()
	//controllers.Trigger()

}
