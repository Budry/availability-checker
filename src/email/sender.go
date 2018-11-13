package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (message *Message) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", message.From)
	if len(message.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(message.To, ";"))
	}
	header += fmt.Sprintf("Subject: %s\r\n", message.Subject)
	header += "\r\n" + message.Body

	return header
}


func Send(message *Message) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	addr := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	err := smtp.SendMail(addr, auth, message.From, message.To, []byte(message.BuildMessage()))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
