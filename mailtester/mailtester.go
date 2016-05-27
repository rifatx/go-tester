package mailtester

import (
	"log"
	"net/smtp"
	"os"
)

func Test() {
	hostname := "mail.smtp2go.com"
	log.SetOutput(os.Stdout)

	auth := smtp.PlainAuth("", "rifatx", "osman.666", hostname)

	to := []string{"rifat.aricanli@gmail.com"}

	msg := []byte("Subject: hey what's up!!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail(
		hostname+":2525",
		auth,
		"sender@example.org",
		to,
		msg,
	)

	if err != nil {
		log.Fatal(err)
	}
}
