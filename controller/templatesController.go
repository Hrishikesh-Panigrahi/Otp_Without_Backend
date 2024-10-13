package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Emailhandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func OTPhandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/otp.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}
