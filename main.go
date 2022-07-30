package main

import (
	"log"
	"net/http"

	//"reflect"
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

	log.Fatal(http.ListenAndServe(":8080", router))
	//wg.Done()
}
func Mailer(wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		println("waiting for trigger")
		controllers.Trigger()

	}
	//println("waiting for trigger")

	// println(reflect.TypeOf(from1))
	// println(from, pass, to, host, port)
	//controllers.SendMail("jesalkrypto@zohomail.in", "Qwertyuiop@1234", "jesalpatel290@gmail.com", "smtp.zoho.in", "587", "msg")
	//controllers.SendMail2(from1, pass, to, host, port, "THIS IS TEST")
	//wg.Done()
	//controllers.Trigger()

}
func main() {

	//run database
	configs.ConnectDB()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go theRouter(wg)
	go Mailer(wg)
	//controllers.SendMail2()
	//controllers.SendMail("jesalkrypto@zohomail.in", "Qwertyuiop@1234", "jesalpatel290@gmail.com", "smtp.zoho.in", "587", "msg")
	wg.Wait()
	//controllers.Trigger()

}
