package mail

import (
	"fmt"
	"testing"
)

func TestSendPlainText(t *testing.T) {
	mail := Setup()
	fmt.Println(mail.To("jinzhu@example.com").From("jinzhu@from.com").Subject("subject").Body("body").String())
}
