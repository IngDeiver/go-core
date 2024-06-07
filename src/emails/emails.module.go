package emails

import (
	email "github.com/ingdeiver/go-core/src/emails/application/services"
	smtp "github.com/ingdeiver/go-core/src/emails/infrastructure/gomail"
)

var SmtpService *smtp.Gomail
var EmailService *email.EmailService

// Instance Repositories, Services, Controllers and more
func InitEmailsModule() {
	// ----- services -----
	SmtpService = smtp.New()
	EmailService = email.New(SmtpService)
}