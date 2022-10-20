package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/9glt/go-sconv"
	"gopkg.in/gomail.v2"
)

func main() {
	if os.Getenv("MAIL_FROM") == "" || os.Getenv("MAIL_USERNAME") == "" || os.Getenv("MAIL_PASSWORD") == "" || os.Getenv("MAIL_HOST") == "" {
		fmt.Printf("please define environement variable\n 'MAIL_FROM'\n 'MAIL_USERNAME'\n 'MAIL_PASSWORD'\n 'MAIL_HOST'\n 'MAIL_PORT'\n")
		log.Fatal("MAIL environment variables are not set")

	}
	if os.Getenv("MAIL_PORT") == "" {
		os.Setenv("MAIL_PORT", "25")
	}
	if os.Getenv("SERVER_BINDTO") == "" {
		os.Setenv("SERVER_BINDTO", "127.0.0.1:8888")
	}

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading body: %v", err)
			io.WriteString(w, "Error reading request body")
			return
		}
		if r.FormValue("to") == "" || r.FormValue("subject") == "" {
			log.Printf("error send parametre missing")
			io.WriteString(w, "Error sending email, to or subject are empty")
			return
		}
		go sendMail(r.FormValue("to"), r.FormValue("subject"), string(body))
	})

	log.Printf("Starting server on: %v", os.Getenv("SERVER_BINDTO"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_BINDTO"), nil))
}

func sendMail(to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(os.Getenv("MAIL_HOST"), sconv.String(os.Getenv("MAIL_PORT")).Int(), os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		log.Printf("error while sending email: %v", err)
	}

}
