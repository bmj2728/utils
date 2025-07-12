package strutil

type PhoneFormat int

const (
	PhoneUS PhoneFormat = iota
	PhoneInternational
	PhoneDigitsOnly
)
