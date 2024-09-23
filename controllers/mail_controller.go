package controllers

import (
	"encoding/json"
	"fmt"
	"mailer-microservice/models"
	"mailer-microservice/services"
	"net/http"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	var mail models.Mail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.SendMail(mail)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}
