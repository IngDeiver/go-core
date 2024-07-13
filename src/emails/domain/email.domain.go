package emailDomain

import "io"

type EmailMessageDomain struct {
	To []string
	Cc *EmailCCDomain
	Subject string
	Attachments []EmailAttachment
}

type EmailAttachment struct {
	Filename string
	Content  io.Reader
}