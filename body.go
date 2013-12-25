package mail

import "fmt"

type Body struct {
	Value       string
	ContentType string
	mailer      *Mailer
}

func (b Body) contentType() string {
	if len(b.ContentType) > 0 {
		return "Content-Type: " + b.ContentType
	}
	return "Content-Type: text/plain; charset=UTF-8"
}

func (b Body) contentTransferEncoding() string {
	return fmt.Sprintf("Content-Transfer-Encoding: %s", "base64")
}

func (b Body) Encode() string {
	return fmt.Sprintf("%v\r\r\n%v\r\n\r\n%v\r\n", b.contentType(), b.contentTransferEncoding(), b.Value)
}
