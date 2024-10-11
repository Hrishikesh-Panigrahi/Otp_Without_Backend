package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")
	t.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/userinput", controller.UserInput)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
