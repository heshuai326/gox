package mobile

import (
	"fmt"
	"github.com/gopub/gox"
	"strings"
)

type PhoneNumber gox.PhoneNumber

func (pn *PhoneNumber) String() string {
	return (*gox.PhoneNumber)(pn).String()
}

func (pn *PhoneNumber) InternationalFormat() string {
	return (*gox.PhoneNumber)(pn).InternationalFormat()
}

type PhoneNumberList struct {
	List []*gox.PhoneNumber
}

func (l *PhoneNumberList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *PhoneNumberList) Get(index int) *PhoneNumber {
	return (*PhoneNumber)(l.List[index])
}

func NewPhoneNumberList() *PhoneNumberList {
	return &PhoneNumberList{}
}

func (l *PhoneNumberList) Add(phoneNumber *PhoneNumber) {
	l.List = append(l.List, (*gox.PhoneNumber)(phoneNumber))
}

func (l *PhoneNumberList) Contains(phoneNumber *PhoneNumber) bool {
	for _, pn := range l.List {
		if pn.String() == (*gox.PhoneNumber)(phoneNumber).String() {
			return true
		}
	}
	return false
}

func (l *PhoneNumberList) ContainsString(phoneNumber string) bool {
	for _, pn := range l.List {
		if pn.String() == phoneNumber {
			return true
		}
	}
	return false
}

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	return (*PhoneNumber)(gox.NewPhoneNumber(callingCode, number))
}

func ParsePhoneNumber(s string, code int) *PhoneNumber {
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	if len(s) == 0 {
		return nil
	}
	pn, _ := gox.ParsePhoneNumber(s)
	if pn != nil {
		return (*PhoneNumber)(pn)
	}
	s = fmt.Sprintf("+%d%s", code, s)
	pn, _ = gox.ParsePhoneNumber(s)
	return (*PhoneNumber)(pn)
}

func TidyPhoneNumber(s string, code int) *PhoneNumber {
	return (*PhoneNumber)(gox.TidyPhoneNumber(s, code))
}
