package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... darkness my old friend")
	fmt.Println("Endpoint Hit: Main")
}

func returnMailSetting(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is mail setting")
	fmt.Println("Endpoint Hit: Mail Setting")
}

func getMailRecipient(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is get mail recipient")
	fmt.Println("Endpoint Hit: Set Mail Recipient")
}

func getMailMessage(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is get mail message")
	fmt.Println("Endpoint Hit: Set Mail Message")
}

func setMailRecipient(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is set mail recipient")
	fmt.Println("Endpoint Hit: Set Mail Recipient")
}

func setMailMessage(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is set mail message")
	fmt.Println("Endpoint Hit: Set Mail Message")
}

func startServer(port string) {
	router := httprouter.New()

	router.GET("/", home)
	router.GET("/mail-setting", returnMailSetting)
	router.GET("/mail-recipient", setMailRecipient)
	router.GET("/mail-message", setMailMessage)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	port := strconv.Itoa(4242)
	fmt.Println("Starting zergver at " + port)
	startServer(port)
}
