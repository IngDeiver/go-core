package gomail

import (
	"bytes"
	"os"
	"strconv"
	"time"

	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	emailDomain "github.com/ingdeiver/go-core/src/emails/domain"
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
	emailHelpers "github.com/ingdeiver/go-core/src/emails/infrastructure"
	"gopkg.in/gomail.v2"
)

var l = logger.Get()

func createGomailDialer() *gomail.Dialer{
	emailEnvs := emailHelpers.GetEmailEnvs()
	port, err := strconv.Atoi(emailEnvs["SMTP_PORT"])
	if err != nil {
		l.Fatal().Msg(err.Error())
	}


	d := gomail.NewDialer(emailEnvs["SMTP_HOST"], port, emailEnvs["SMTP_USER"], emailEnvs["SMTP_PASWORD"])
	return d
}

func getGomailMessage(emailInfo emailDomain.EmailMessageDomain) *gomail.Message {
	m := gomail.NewMessage()
	from := os.Getenv("EMAIL_FROM")

	if len(from) == 0{
		l.Fatal().Msg("email 'from' env missing")
	}

	m.SetHeader("From", from)
	m.SetHeader("To", emailInfo.To...)

	if emailInfo.Cc != nil{
		m.SetAddressHeader("Cc", emailInfo.Cc.Address, emailInfo.Cc.Name)
	}

	m.SetHeader("Subject",emailInfo.Subject)
	return m
	
}


var ch chan *emailDomain.EmailChanel
// implements SMTPServiceDomain
type Gomail struct {

}

func New() *Gomail{
	return &Gomail{}
}

func (Gomail) SendEmail(emailType emailConstants.EmailType, 
	emailInfo emailDomain.EmailMessageDomain, 
	templateInfo emailDomain.EmailTemplateBodyDomain ) (bool, error){
	m:= getGomailMessage(emailInfo)
	template := emailHelpers.LoadEmailTemplateByType(emailType)
	var body bytes.Buffer
	err := template.Execute(&body, templateInfo)

	if err != nil {
		l.Error().Msgf("Error executing template: %v", err.Error())
		return false, err
	}
	m.SetBody("text/html", body.String())


	d := createGomailDialer()

	if err := d.DialAndSend(m); err != nil {
		l.Error().Msgf("Error to send email: %v", err.Error())
	}
	l.Info().Msgf("Email sent to %v: ", emailInfo.To)
	return true, nil
	
}

func (Gomail) CreateEmailsDeamon() chan *emailDomain.EmailChanel {
	if ch != nil {
		return ch
	}

	ch := make(chan *emailDomain.EmailChanel)

	go func() {
		d := createGomailDialer()
	
		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case data, ok := <-ch:
				if !ok { // chanal is closed
					l.Warn().Msg("The email chanel is closed, mail omitted")
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}

				m:= getGomailMessage(data.Message)
				template := emailHelpers.LoadEmailTemplateByType(data.EmailType)
				var body bytes.Buffer
				err := template.Execute(&body, data.TemplateBody)

				if err != nil {
					l.Error().Msgf("Error executing template: %v", err.Error())
					return
				}
				m.SetBody("text/html", body.String())
				if err := gomail.Send(s, m); err != nil {
					l.Err(err)
				}
				l.Info().Msgf("Email sent to %v: ", data.Message.To)
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()
	l.Info().Msg("Emails channel started")
	// Use the channel in your program to send emails and close it
	return ch
}