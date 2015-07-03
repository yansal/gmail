package main

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/howeyc/gopass"
)

func main() {
	fmt.Print("Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Password:")
	pass := gopass.GetPasswd()

	fmt.Print("To: ")
	var to string
	fmt.Scan(&to)

	auth := smtp.PlainAuth("", username, string(pass), "smtp.gmail.com")
	msg := []byte("Subject: This is the subject\nThis is the body")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "", []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}
