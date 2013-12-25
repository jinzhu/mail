package mail

type Body struct {
	Value   string
	Charset string
	mailer  *Mailer
}

func (b Body) Encode() string {
	return "\r\n" + b.Value + "\r\n"
}
