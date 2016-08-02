package gmail

import "net/smtp"

var host = "smtp.gmail.com"

// SendMail sends an email with a Gmail account
func SendMail(from, password string, to []string, msg []byte) error {
	auth := smtp.PlainAuth("", from, password, host)
	return smtp.SendMail(host+":587", auth, from, to, msg)
}
