package emailService

import (
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	emailDomain "github.com/ingdeiver/go-core/src/emails/domain"
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
	emailSmtpDomain "github.com/ingdeiver/go-core/src/emails/domain/interfaces"
)

var l = logger.Get()

type EmailService struct {
	smtpService emailSmtpDomain.SMTPServiceDomain
	Channel chan *emailDomain.EmailChanel
}

func New(smtpService emailSmtpDomain.SMTPServiceDomain) *EmailService{
	channel := smtpService.CreateEmailsDeamon()
	return &EmailService{smtpService, channel}
}

func (s *EmailService) SendEmail(emailType emailConstants.EmailType, 
	emailInfo emailDomain.EmailMessageDomain, 
	templateInfo emailDomain.EmailTemplateBodyDomain ) (bool, error){
	return s.smtpService.SendEmail(emailType,emailInfo,templateInfo)
}


func (s *EmailService) AddEmailToChannel(emailType emailConstants.EmailType, 
	emailInfo emailDomain.EmailMessageDomain, 
	templateInfo emailDomain.EmailTemplateBodyDomain){
		if s.Channel == nil {
			l.Error().Msg("Nt found emails channel")
		}
		s.Channel <- &emailDomain.EmailChanel{EmailType: emailType, Message: emailInfo, TemplateBody: templateInfo }
}