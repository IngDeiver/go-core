package smtpServiceDomain

import (
	emailDomain "github.com/ingdeiver/go-core/src/emails/domain"
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
)

type SMTPServiceDomain interface {
	SendEmail(emailType emailConstants.EmailType, emailInfo emailDomain.EmailMessageDomain, template emailDomain.EmailTemplateBodyDomain ) (bool, error)
	CreateEmailsDeamon() chan *emailDomain.EmailChanel
}