package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	gomail "gopkg.in/gomail.v2"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.POST("/feedback", feedback)

	e.Logger.Fatal(e.Start(":8083"))
}

func feedback(c echo.Context) error {
	body := c.FormValue("body")
	subject := c.FormValue("subject")
	fmt.Println(body, subject)
	Email(body, subject)

	return c.JSON(http.StatusOK, "ok")
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func Email(body, subject string) {
	from := "spektr.feedback@mail.ru"
	pass := os.Getenv("MAIL_PASS")
	to := "spektr.feedback@mail.ru"
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	n := gomail.NewDialer("smtp.mail.ru", 587, from, pass)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}
