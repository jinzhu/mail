package mail

import (
	"encoding/base64"
	"fmt"
)

type Attachment struct {
	FileName    string
	Content     []byte
	ContentType string
	Inline      bool
}

func (s Attachment) contentType() string {
	if len(s.ContentType) > 0 {
		return "Content-Type: " + s.ContentType
	}
	return "Content-Type: application/octet-stream"
}

func (s Attachment) contentTransferEncoding() string {
	return fmt.Sprintf("Content-Transfer-Encoding: %s", "base64")
}

func (s Attachment) contentDisposition() string {
	return fmt.Sprintf(`Content-Disposition: attachment; filename="%v"`, s.FileName)
}

func (s *Attachment) URL() string {
	return ""
}

func (s *Attachment) Encode() string {
	return fmt.Sprintf("%v\r\r\n%v\r\n%v\r\n\r\n%v\r\n",
		s.contentType(), s.contentTransferEncoding(), s.contentDisposition(), base64.StdEncoding.EncodeToString(s.Content))
}
