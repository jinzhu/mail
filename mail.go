package mail

type Mail struct {
	Charset                 string
	ContentTransferEncoding string
	Error                   error
	From                    string
	To                      []string
	Cc                      []string
	Bcc                     []string
	Subject                 string
	Bodys                   []Body
	Headers                 []Header
	Attachments             []Attachment
}

func (s Mail) clone() Mail {
	return Mail{
		Charset:                 s.Charset,
		ContentTransferEncoding: s.ContentTransferEncoding,
		Error:       s.Error,
		From:        s.From,
		To:          s.To,
		Cc:          s.Cc,
		Bcc:         s.Bcc,
		Subject:     s.Subject,
		Bodys:       s.Bodys,
		Headers:     s.Headers,
		Attachments: s.Attachments,
	}
}
