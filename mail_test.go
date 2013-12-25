package mail

import (
	"fmt"
	"testing"
)

func TestSendPlainText(t *testing.T) {
	mail := Setup()
	fmt.Println(mail.To("jinzhu@example.com").From("jinzhu@from.com").Subject("subject").
		Body("text").Body(Body{Value: "html", ContentType: "text/html; charset=UTF-8"}).
		Attach("/home/jinzhu/ffff").
		String())

	fmt.Println(mail.To("jinzhu@example.com").From("jinzhu@from.com").Subject("subject").Body("text").String())
}
