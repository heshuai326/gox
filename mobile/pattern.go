package mobile

import (
	"github.com/gopub/gox"
)

func IsEmail(email string) bool {
	return gox.IsEmail(email)
}

func IsLink(s string) bool {
	return gox.IsLink(s)
}

func IsName(name string) bool {
	return gox.IsName(name)
}

func IsPhoneNumber(phoneNumber string) bool {
	return gox.IsPhoneNumber(phoneNumber)
}

func IsBirthDate(s string) bool {
	return gox.IsBirthDate(s)
}
