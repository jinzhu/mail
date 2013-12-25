package mail

import (
	"fmt"
	"strings"
	"time"
)

type Mailer struct {
	Mail
	Boundary string
}

func (s *Mailer) clone() *Mailer {
	return &Mailer{s.Mail.clone()}
}

func (m *Mailer) boundary() string {
	if len(m.Boundary) == 0 {
		m.Boundary = fmt.Sprintf("_mimepart_%v", time.Now().UnixNano())
	}
	return m.Boundary
}

func (m *Mailer) crlfBoundary() string {
	if len(m.contentType()) == 0 {
		return ""
	}
	return fmt.Sprintf("\r\n--%v\r\n", m.boundary())
}

func (m *Mailer) endBoundary() string {
	if len(m.contentType()) == 0 {
		return ""
	}
	return fmt.Sprintf("\r\n--%v--\r\n", m.boundary())
}

func (m *Mailer) charset() string {
	if len(m.Mail.Charset) > 0 {
		return m.Mail.Charset
	} else {
		return "utf-8"
	}
}

func (m *Mailer) contentType() string {
	var content_type string
	if len(m.Mail.Attachments) > 0 || len(m.Mail.Bodys) > 1 {
		content_type = fmt.Sprintf(`Content-Type: multipart/alternative; boundary="%v"`, m.boundary())
	} else {
		return ""
	}

	return fmt.Sprintf(`%v; charset="%v"`, content_type, m.charset())
}

func (m *Mailer) contentTransferEncoding() string {
	if len(m.contentType()) == 0 {
		return ""
	}
	return "base64"
}

func (m *Mailer) String() (string, error) {
	message, mail, header := "", m.Mail, make(map[string]string)

	header["From"] = mail.From
	header["To"] = strings.Join(mail.To, ",")
	header["Subject"] = mail.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = m.contentType()
	header["Content-Transfer-Encoding"] = m.contentTransferEncoding()

	for k, v := range header {
		if len(v) > 0 {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
	}

	if len(m.Mail.Bodys) == 0 {
		m.Body("")
	}

	for _, body := range m.Mail.Bodys {
		message += m.crlfBoundary()
		message += body.Encode() + "\r\n"
	}

	for _, attachment := range m.Mail.Attachments {
		message += m.crlfBoundary()
		message += attachment.Encode() + "\r\n"
	}

	message += m.endBoundary()

	return message, nil
}
