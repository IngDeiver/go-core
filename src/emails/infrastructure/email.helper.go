package emailHelpers

import (
	"embed"
	"html/template"
	"os"

	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	emailConstants "github.com/ingdeiver/go-core/src/emails/domain/constants"
)

var l = logger.Get()

//go:embed templates/password.html
var passwordTemplateFS embed.FS

//go:embed templates/notification.html
var notificationTemplateFS embed.FS

func LoadEmailTemplateByType(emailType emailConstants.EmailType) *template.Template{
  var tmp *template.Template
	switch  emailType {
		case emailConstants.Password:
			htmlTemplate,err := template.ParseFS(passwordTemplateFS, "templates/password.html")

			if err != nil {
				l.Fatal().Msg(err.Error())
			}
			tmp = htmlTemplate
		case emailConstants.Notification:
			htmlTemplate,err := template.ParseFS(notificationTemplateFS, "templates/notification.html")

			if err != nil {
				l.Fatal().Msg(err.Error())
			}
			tmp = htmlTemplate
	}
	return tmp
}

func GetEmailEnvs() (map[string]string) {
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	SMTP_USER := os.Getenv("SMTP_USER")
	SMTP_PASWORD := os.Getenv("SMTP_PASWORD")
	if len(SMTP_HOST) == 0 || len(SMTP_PORT) == 0  || len(SMTP_USER) == 0 || len(SMTP_PASWORD) == 0 {
		l.Fatal().Msg("No found email enviroments")
	}

	return map[string]string{
		"SMTP_HOST":SMTP_HOST,
		"SMTP_PORT":SMTP_PORT,
		"SMTP_USER":SMTP_USER,
		"SMTP_PASWORD":SMTP_PASWORD,
	}
}