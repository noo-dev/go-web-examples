package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Subject string
	Email   string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		fmt.Printf("Email: %v \n Subject: %v \n Message: %v", details.Email, details.Subject, details.Message)

		tmpl.Execute(w, struct{ Success bool }{true})

	})

	http.ListenAndServe(":80", nil)
}
