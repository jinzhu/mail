package mail

type Attachment struct {
	FileName string
	MimeType string
	Content  string
	Encoding string
	Inline   bool
}

func (attacchment *Attachment) URL() string {
	return ""
}

func (attachment *Attachment) Encode() string {
	return ""
}
