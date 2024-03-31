package emailService

import (
	emailDomain "github.com/ingdeiver/go-core/src/emails/domain"
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
	emailSmtpDomain "github.com/ingdeiver/go-core/src/emails/domain/interfaces"
)

type EmailService struct {
	smtpService emailSmtpDomain.SMTPServiceDomain
}

func New(smtpService emailSmtpDomain.SMTPServiceDomain) EmailService{
	return EmailService{smtpService}
}

func (s *EmailService) SendEmail(emailType emailConstants.EmailType, 
	emailInfo emailDomain.EmailMessageDomain, 
	templateInfo emailDomain.EmailTemplateBodyDomain ) (bool, error){
	return s.smtpService.SendEmail(emailType,emailInfo,templateInfo)
}

func (s *EmailService) CreateEmailsDeamon() chan *emailDomain.EmailChanel {
	return s.smtpService.CreateEmailsDeamon()
}