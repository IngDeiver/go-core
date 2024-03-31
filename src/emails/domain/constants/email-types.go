package emailTypes

type EmailType int

const (
    Password EmailType = iota
    Notification
)
