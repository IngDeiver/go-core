package emailDomain

import (
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
)
type EmailChanel struct {
	EmailType emailConstants.EmailType
	Message EmailMessageDomain
	TemplateBody EmailTemplateBodyDomain
}