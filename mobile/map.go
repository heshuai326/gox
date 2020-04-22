package mobile

import (
	"github.com/gopub/gox"
	"github.com/gopub/gox/geo"
)

type Map struct {
	m gox.M
}

func NewMap() *Map {
	return &Map{m: gox.M{}}
}

func (m *Map) GetInt64(key string) int64 {
	return m.m.Int64(key)
}

func (m *Map) GetFloat64(key string) float64 {
	return m.m.Float64(key)
}

func (m *Map) GetString(key string) string {
	return m.m.String(key)
}

func (m *Map) GetBool(key string) bool {
	return m.m.Bool(key)
}

func (m *Map) GetInt64List(key string) *Int64List {
	switch v := m.m[key].(type) {
	case []int64:
		return &Int64List{List: v}
	case *Int64List:
		return v
	default:
		return nil
	}
}

func (m *Map) GetStringList(key string) *StringList {
	switch v := m.m[key].(type) {
	case []string:
		return &StringList{List: v}
	case *StringList:
		return v
	default:
		return nil
	}
}

func (m *Map) GetPhoneNumber(key string) *gox.PhoneNumber {
	return m.m.PhoneNumber(key)
}

func (m *Map) GetLocation(key string) *geo.Point {
	v, _ := m.m[key].(*geo.Point)
	return v
}

func (m *Map) SetInt64(key string, val int64) {
	m.m[key] = val
}

func (m *Map) SetFloat64(key string, val float64) {
	m.m[key] = val
}

func (m *Map) SetString(key string, val string) {
	m.m[key] = val
}

func (m *Map) SetBool(key string, val bool) {
	m.m[key] = val
}

func (m *Map) SetPhoneNumber(key string, val *gox.PhoneNumber) {
	m.m[key] = val
}

func (m *Map) SetInt64s(key string, val *Int64List) {
	m.m[key] = val.List
}

func (m *Map) SetStrings(key string, val *StringList) {
	m.m[key] = val.List
}

func (m *Map) SetLocation(key string, val *geo.Point) {
	m.m[key] = val
}

func (m *Map) M() gox.M {
	return m.m
}
