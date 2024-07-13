package emails

import (
	email "github.com/ingdeiver/go-core/src/emails/application/services"
	gomail "github.com/ingdeiver/go-core/src/emails/infrastructure/gomail"
)

var SmtpService *gomail.Gomail
var EmailService *email.EmailService

// Instance Repositories, Services, Controllers and more
func InitEmailsModule() {
	// ----- services -----
	SmtpService = gomail.New()
	EmailService = email.New(SmtpService)
}