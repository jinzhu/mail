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
	new_mail := m.clone()
	new_mail.Mail.From = value
	return new_mail
}

func (m *Mailer) To(values ...string) *Mailer {
	new_mail := m.clone()
	new_mail.Mail.To = append(new_mail.Mail.To, values...)
	return new_mail
}

func (m *Mailer) Cc(values ...string) *Mailer {
	new_mail := m.clone()
	new_mail.Mail.Cc = append(new_mail.Mail.Cc, values...)
	return new_mail
}

func (m *Mailer) Bcc(values ...string) *Mailer {
	new_mail := m.clone()
	new_mail.Mail.Bcc = append(new_mail.Mail.Bcc, values...)
	return new_mail
}

func (m *Mailer) Subject(value string) *Mailer {
	new_mail := m.clone()
	new_mail.Mail.Subject = value
	return new_mail
}

func (m *Mailer) Body(values ...interface{}) *Mailer {
	new_mail := m.clone()
	for _, value := range values {
		if str, ok := value.(string); ok {
			new_mail.Mail.Bodys = append(m.Mail.Bodys, Body{Value: str, mailer: m})
		} else if body, ok := value.(Body); ok {
			body.mailer = m
			new_mail.Mail.Bodys = append(m.Mail.Bodys, body)
		} else {
			new_mail.Mail.Error = errors.New(fmt.Sprint("Unknown body value", value))
		}
	}
	return new_mail
}

func (m *Mailer) Header(key, value string) *Mailer {
	new_mail := m.clone()
	header := Header{Key: key, Value: value}
	new_mail.Mail.Headers = append(m.Mail.Headers, header)
	return new_mail
}

func (m *Mailer) Attach(attachment interface{}) *Mailer {
	new_mail := m.clone()
	if filename, ok := attachment.(string); ok {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			new_mail.Mail.Error = err
		}
		attachment = Attachment{Content: b, FileName: filepath.Base(filename)}
	}
	if attach, ok := attachment.(Attachment); ok {
		new_mail.Mail.Attachments = append(m.Mail.Attachments, attach)
	}
	return new_mail
}

func (m *Mailer) Charset(str string) *Mailer {
	new_mail := m.clone()
	new_mail.Mail.Charset = str
	return new_mail
}

func (m *Mailer) Send() error {
	message, err := m.String()
	if err == nil {
		fmt.Println(message)
	}
	return err
}
