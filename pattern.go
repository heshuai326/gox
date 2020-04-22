package gox

import (
	"net/url"
	"regexp"
	"time"

	"github.com/gopub/log"
	"github.com/nyaruka/phonenumbers"
)

var nameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-._]*$`)
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	return emailRegexp.MatchString(email)
}

func IsLink(s string) bool {
	if len(s) == 0 {
		return false
	}
	u, err := url.Parse(s)
	if err != nil {
		return false
	}
	if len(u.Scheme) == 0 {
		return false
	}

	if len(u.Host) == 0 {
		return false
	}
	return true
}

func IsName(name string) bool {
	if len(name) == 0 {
		return false
	}
	return nameRegexp.MatchString(name)
}

func IsPhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) == 0 {
		return false
	}
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		log.Error(err)
		return false
	}

	return phonenumbers.IsValidNumber(parsedNumber)
}

func IsBirthDate(s string) bool {
	t, err := time.Parse("2006-1-2", s)
	if err != nil {
		log.Error(err)
		return false
	}

	if t.After(time.Now()) {
		return false
	}

	return true
}
