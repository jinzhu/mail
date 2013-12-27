package mail

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
)

type Mailer struct {
	Mail
	Boundary string
}

func (s *Mailer) clone() *Mailer {
	return &Mailer{Mail: s.Mail.clone()}
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
	message, header := "", make(map[string]string)

	headers := map[string][]string{
		"From": []string{m.Mail.From},
		"To":   m.Mail.To,
		"Cc":   m.Mail.Cc,
		"Bcc":  m.Mail.Bcc,
	}

	for key, addresses := range headers {
		formated_addresses := []string{}
		for _, address := range addresses {
			parsed_addresses, err := mail.ParseAddressList(address)
			if err == nil {
				for _, parsed_address := range parsed_addresses {
					if len(parsed_address.Name) > 0 {
						formated_addresses = append(formated_addresses, fmt.Sprintf("%v <%v>", parsed_address.Name, parsed_address.Address))
					} else {
						formated_addresses = append(formated_addresses, fmt.Sprintf("%v", parsed_address.Address))
					}
				}
			}
		}
		header[key] = strings.Join(formated_addresses, ", ")
	}

	header["Subject"] = m.Mail.Subject
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
