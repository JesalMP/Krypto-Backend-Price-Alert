package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	//"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	//"net/smtp"
	"github.com/JesalMP/Krypto-Backend-Price-Alert/configs"
	"github.com/JesalMP/Krypto-Backend-Price-Alert/models"
	"github.com/JesalMP/Krypto-Backend-Price-Alert/responses"
	gomail "gopkg.in/gomail.v2"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var alertCollection *mongo.Collection = configs.GetCollection(configs.DB, "Alerts")
var validate = validator.New()

func CreateAlert() http.HandlerFunc {
	//print("sss")
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//var alert models.Alert
		params := mux.Vars(r)
		pricealert := params["alertPrice"]
		//i, err := strconv.Atoi(pricealert)

		i, err := strconv.ParseFloat(pricealert, 64)
		if err != nil {
			// ... handle error
			panic(err)
		}
		defer cancel()

		//validate the request body

		newAlert := models.Alert{
			Id:                   primitive.NewObjectID(),
			AlertPrice:           i,
			PriceAtAlertCreation: GetPrice(),
			AlertState:           "Primed",
			AlertCreationTime:    time.Now().String(),
			AlertTriggerTime:     "Not Yet Triggered",
		}
		result, err := alertCollection.InsertOne(ctx, newAlert)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}
func GetAlertsState() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		state := params["state"]
		var alert []models.Alert
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		results, err := alertCollection.Find(ctx, bson.M{"alertstate": state})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAlert models.Alert
			if err = results.Decode(&singleAlert); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}
			alert = append(alert, singleAlert)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": alert}}
		json.NewEncoder(rw).Encode(response)
	}
}

func DeleteAlert() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		pricealert := params["alertPrice"]
		i, err := strconv.Atoi(pricealert)
		if err != nil {
			// ... handle error
			panic(err)
		}

		//fmt.Println(pricealert, i)
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := alertCollection.DeleteMany(ctx, bson.M{"alertprice": i})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllAlerts() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.Alert
		defer cancel()

		results, err := alertCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAlert models.Alert
			if err = results.Decode(&singleAlert); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}
			users = append(users, singleAlert)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetPrice() (f float64) {
	response, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	x := map[string]string{}

	json.Unmarshal([]byte(string(responseData)), &x)
	//fmt.Println(x["price"])
	f1, err := strconv.ParseFloat(x["price"], 64)
	f = f1
	//println(f)
	return
}
func SendMail(from1, pass, to1, host, port, msg string) {
	from := from1
	password := pass

	// Receiver email address.
	to := []string{
		to1,
	}

	// smtp server configuration.
	smtpHost := host
	smtpPort := port

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your Trigger has Triggered \n%s\n\n", mimeHeaders)))
	//numi := models.EmailBody{"Jesal @ Krypto Employee", msg}
	//t.Execute(body, numi)

	// Sending email.
	println(body.Bytes())
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

}
func SendMail2(from1, pass, to1, host, port, msg1 string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from1)
	msg.SetHeader("To", to1)
	msg.SetHeader("Subject", "BTC Price Alert!!!")
	msg.SetBody("text/html", msg1)
	//msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer(host, 587, from1, pass)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	fmt.Println("Email Sent!")
}
func Trigger() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var alerts []models.Alert
	defer cancel()

	//objId, _ := primitive.ObjectIDFromHex(userId)

	results, err := alertCollection.Find(ctx, bson.M{"alertstate": "Primed"})
	if err != nil {

		panic(err)
	}
	var currPrice float64 = GetPrice()
	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAlert models.Alert
		if err = results.Decode(&singleAlert); err != nil {
			panic(err)
		}
		alerts = append(alerts, singleAlert)
		//print(alert)
	}
	from := string(configs.EnvEmailMailHandler())
	//from1 := from
	pass := string(configs.EnvEmailPassHandler())
	to := string(configs.EnvUserEmailHandler())
	//toList := []string{to}
	host := string(configs.EnvHostHandler())
	port := string(configs.EnvPortHandler())
	//println(reflect.TypeOf(from1))
	//println(from, pass, to, host, port)
	//&& alert.AlertPrice >= alert.PriceAtAlertCreation
	for _, alert := range alerts {
		if currPrice >= alert.AlertPrice && alert.AlertPrice >= alert.PriceAtAlertCreation {
			var msg string
			msg = "This message is by Krypto\n Your Trigger of value " + fmt.Sprintf("%v", alert.AlertPrice) + " On Bitcoin Prices is Triggered as of " + time.Now().String() + ", Check Binance for more detailed prices. Current Price : " + fmt.Sprintf("%v", currPrice)
			SendMail2(from, pass, to, host, port, msg)
			//SendMail("jesalkrypto@zohomail.in", "Qwertyuiop@1234", "jesalpatel290@gmail.com", "smtp.zoho.in", "587", "msg")
			idcurr := alert.Id
			//println(idcurr.Hex())
			// //docID := "5d1719988f83df290e8c92ca"
			objID, err := primitive.ObjectIDFromHex(idcurr.Hex())
			if err != nil {
				panic(err)
			}
			//filter := bson.M{"_id": bson.M{"$eq": objID}}
			update := bson.M{"$set": bson.M{"alertstate": "Triggered", "alerttriggertime": time.Now().String()}}
			// result, err := alertCollection.UpdateOne(
			// 	context.Background(),
			// 	filter,
			// 	update,
			// )
			result, err := alertCollection.UpdateOne(ctx, bson.M{"id": objID}, update)
			//print(result)
			if err != nil {
				panic(err)
				print(result)
			}
			// println(objID)

		}
		if currPrice <= alert.AlertPrice && alert.AlertPrice <= alert.PriceAtAlertCreation {
			var msg string
			msg = "This message is by Krypto\n Your Trigger of value " + fmt.Sprintf("%v", alert.AlertPrice) + " On Bitcoin Prices is Triggered as of " + time.Now().String() + ", Check Binance for more detailed prices. Current Price : " + fmt.Sprintf("%v", currPrice)
			SendMail(from, pass, to, host, port, msg)
			idcurr := alert.Id
			//println(idcurr.Hex())
			// //docID := "5d1719988f83df290e8c92ca"
			objID, err := primitive.ObjectIDFromHex(idcurr.Hex())
			if err != nil {
				panic(err)
			}
			//filter := bson.M{"_id": bson.M{"$eq": objID}}
			update := bson.M{"$set": bson.M{"alertstate": "Triggered", "alerttriggertime": time.Now().String()}}
			// result, err := alertCollection.UpdateOne(
			// 	context.Background(),
			// 	filter,
			// 	update,
			// )
			result, err := alertCollection.UpdateOne(ctx, bson.M{"id": objID}, update)
			//print(result)
			if err != nil {
				panic(err)
				print(result)
			}
		}
	}

}
