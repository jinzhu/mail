# Mail

mail.SetupSMTP()
// location: '/usr/sbin/sendmail', arguments: '-i -t'
// address:              'smtp.gmail.com', port:                 587, domain:               'example.com', user_name:            '<username>', password:             '<password>', authentication:       'plain', enable_starttls_auto: true
// preview
// drop

mail.To("wosmvp@gmail.com").From("shop@lacoste.jp").Subject("hello").Body("test", {{template}}).Params(interface{}).Header()

Header(key, value)
Attachment(filename, file, mime_type: ,encoding: ,content, inline)

Attachment().URL()


err := mail.Send()
err := mail.Preview()


mail.To("wosmvp@gmail.com").From("shop@lacoste.jp").Subject("hello").Body("test", {{template}}).Params(interface{}).Header()
