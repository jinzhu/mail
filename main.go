package mail

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func Setup() *Mailer {
	return &Mailer{}
}

func (m *Mailer) From(value string) *Mailer {
	m.Mail.From = value
	return m
}

func (m *Mailer) To(values ...string) *Mailer {
	m.Mail.To = append(m.Mail.To, values...)
	return m
}

func (m *Mailer) Cc(values ...string) *Mailer {
	m.Mail.Cc = append(m.Mail.Cc, values...)
	return m
}

func (m *Mailer) Bcc(values ...string) *Mailer {
	m.Mail.Bcc = append(m.Mail.Bcc, values...)
	return m
}

func (m *Mailer) Subject(value string) *Mailer {
	m.Mail.Subject = value
	return m
}

func (m *Mailer) Body(values ...interface{}) *Mailer {
	for _, value := range values {
		if str, ok := value.(string); ok {
			m.Mail.Bodys = append(m.Mail.Bodys, Body{Value: str, mailer: m})
		} else if body, ok := value.(Body); ok {
			body.mailer = m
			m.Mail.Bodys = append(m.Mail.Bodys, body)
		} else {
			m.Mail.Error = errors.New(fmt.Sprint("Unknown body value", value))
		}
	}
	return m
}

func (m *Mailer) Header(key, value string) *Mailer {
	header := Header{Key: key, Value: value}
	m.Mail.Headers = append(m.Mail.Headers, header)
	return m
}

func (m *Mailer) Attach(attachment interface{}) *Mailer {
	if filename, ok := attachment.(string); ok {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			m.Mail.Error = err
		}
		attachment = Attachment{Content: b, FileName: filepath.Base(filename)}
	}
	if attach, ok := attachment.(Attachment); ok {
		m.Mail.Attachments = append(m.Mail.Attachments, attach)
	}
	return m
}

func (m *Mailer) Charset(str string) *Mailer {
	m.Mail.Charset = str
	return m
}

func (m *Mailer) Send() error {
	message, err := m.String()
	if err == nil {
		fmt.Println(message)
	}
	return err
}
