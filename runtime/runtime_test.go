package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

type Shape struct {
	area float64
}

func TestMakeZero(t *testing.T) {
	var i = 10
	MakeZero(&i)
	if i != 0 {
		t.FailNow()
	}

	var s = Shape{}
	s.area = 10
	MakeZero(&s)
	ss := Shape{}
	if s != ss {
		t.FailNow()
	}
}

func TestIsNil(t *testing.T) {
	//	var s *Shape = nil
	var i error = nil
	if IsNil(i) == false {
		t.Fail()
	}
}

type Person struct {
	Age       int
	Birthdate string
}

func (p *Person) Error() string {
	return ""
}

func TestRenew(t *testing.T) {
	var p1 *Person
	p2 := &Person{}
	//p1 = p2
	Renew(&p1, p2)
	if p1 == nil {
		t.Fail()
	}
}

type AddressInfo struct {
	City string
}

type User struct {
	Name string
	*AddressInfo
}

type Topic struct {
	*User
	Owner *User
	Title string
}

func TestDeepNew(t *testing.T) {
	topic := DeepNew(reflect.TypeOf(Topic{})).Interface().(*Topic)
	fmt.Println(topic.Name, topic.Title, topic.City, topic.Owner.Name)
}
