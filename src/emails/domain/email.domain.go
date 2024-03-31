package emailDomain

type EmailMessageDomain struct {
	To []string
	Cc *EmailCCDomain
	Subject string
	//Attach string
}