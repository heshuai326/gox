package gox

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
)

var mu sync.RWMutex
var nameToPrototype = map[string]reflect.Type{
	"int":     reflect.TypeOf(int(1)),
	"int8":    reflect.TypeOf(int8(1)),
	"int16":   reflect.TypeOf(int16(1)),
	"int32":   reflect.TypeOf(int32(1)),
	"int64":   reflect.TypeOf(int64(1)),
	"uint":    reflect.TypeOf(int(1)),
	"uint8":   reflect.TypeOf(uint8(1)),
	"uint16":  reflect.TypeOf(uint16(1)),
	"uint32":  reflect.TypeOf(uint32(1)),
	"uint64":  reflect.TypeOf(uint64(1)),
	"float32": reflect.TypeOf(float32(1)),
	"float64": reflect.TypeOf(float64(1)),
	"bool":    reflect.TypeOf(true),
	"string":  reflect.TypeOf(""),
}

type AnyType interface {
	AnyType() string
}

// RegisterAny bind typ with prototype
// E.g.
//		contents.Register("image", &contents.Image{})
func RegisterAny(prototype interface{}) {
	name := GetAnyTypeName(prototype)
	mu.Lock()
	defer mu.Unlock()
	if _, ok := nameToPrototype[name]; ok {
		log.Fatalf("Duplicate name %s", name)
	}

	nameToPrototype[name] = reflect.TypeOf(prototype)
}

func GetAnyTypeName(prototype interface{}) string {
	if a, ok := prototype.(AnyType); ok {
		return a.AnyType()
	}

	p := reflect.TypeOf(prototype)
	for p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return CamelToSnake(p.Name())
}

func getProtoType(typ string) (reflect.Type, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if prototype, ok := nameToPrototype[typ]; ok {
		return prototype, true
	}
	return nil, false
}

var _ sql.Scanner = (*Any)(nil)
var _ driver.Valuer = (*Any)(nil)

type Any struct {
	val     interface{}
	jsonStr string
}

func NewAny(v interface{}) *Any {
	a := &Any{}
	a.SetVal(v)
	return a
}

func (a *Any) Val() interface{} {
	return a.val
}

func (a *Any) SetVal(v interface{}) {
	a.val = v
	a.jsonStr = ""
}

func (a *Any) JSONString() string {
	if len(a.jsonStr) == 0 {
		data, _ := json.Marshal(a)
		a.jsonStr = string(data)
	}
	return a.jsonStr
}

func (a *Any) SetImage(i *Image) {
	a.SetVal(i)
}

func (a *Any) SetAudio(au *Audio) {
	a.SetVal(au)
}

func (a *Any) SetVideo(v *Video) {
	a.SetVal(v)
}

func (a *Any) SetString(s string) {
	a.SetVal(s)
}

func (a *Any) SetInt(i int64) {
	a.SetVal(i)
}

func (a *Any) SetFloat(i float64) {
	a.SetVal(i)
}

func (a *Any) SetFile(f *File) {
	a.SetVal(f)
}

func (a *Any) SetWebPage(wp *WebPage) {
	a.SetVal(wp)
}

func (a *Any) Int() int64 {
	v, _ := a.val.(int64)
	return v
}

func (a *Any) Float() float64 {
	i, _ := a.val.(float64)
	return i
}

func (a *Any) Text() string {
	s, _ := a.val.(string)
	return s
}

func (a *Any) Image() *Image {
	img, _ := a.val.(*Image)
	return img
}

func (a *Any) Video() *Video {
	v, _ := a.val.(*Video)
	return v
}

func (a *Any) Audio() *Audio {
	v, _ := a.val.(*Audio)
	return v
}

func (a *Any) File() *File {
	v, _ := a.val.(*File)
	return v
}

func (a *Any) WebPage() *WebPage {
	v, _ := a.val.(*WebPage)
	return v
}

const (
	keyAnyType = "@t"
	keyAnyVal  = "@v"
)

func (a *Any) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	typ, _ := m[keyAnyType].(string)
	pt, found := getProtoType(typ)
	if !found {
		a.val = m[keyAnyVal]
		if a.val == nil {
			return errors.New("value is empty")
		}

		if GetAnyTypeName(a.val) == typ {
			return nil
		}
		return fmt.Errorf("type doesn't match: %s and %s", typ, GetAnyTypeName(a.val))
	}

	if v, ok := m[keyAnyVal]; ok {
		b, _ = json.Marshal(v)
	}

	var ptrVal = reflect.New(pt)

	for val := ptrVal; val.Kind() == reflect.Ptr && val.CanSet(); val = val.Elem() {
		val.Set(reflect.New(val.Elem().Type()))
	}

	err := json.Unmarshal(b, ptrVal.Interface())
	if err != nil {
		return err
	}
	a.SetVal(ptrVal.Elem().Interface())
	return nil
}

func (a *Any) MarshalJSON() ([]byte, error) {
	if a == nil || a.val == nil {
		return json.Marshal(nil)
	}

	var m = make(map[string]interface{})

	t := reflect.TypeOf(a.val)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct || t.Kind() == reflect.Map {
		b, err := json.Marshal(a.val)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
	} else {
		m[keyAnyVal] = a.val
	}

	m[keyAnyType] = a.TypeName()
	return json.Marshal(m)
}

func (a *Any) TypeName() string {
	return GetAnyTypeName(a.val)
}

func (a *Any) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	if s, ok := src.(string); ok {
		return json.Unmarshal([]byte(s), a)
	} else if b, ok := src.([]byte); ok {
		return json.Unmarshal(b, a)
	} else {
		return fmt.Errorf("invalid type:%v", reflect.TypeOf(src))
	}
}

func (a *Any) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

func init() {
	RegisterAny(&Image{})
	RegisterAny(&Video{})
	RegisterAny(&Audio{})
	RegisterAny(&WebPage{})
	RegisterAny(&File{})
}

type Image struct {
	Link      string `json:"link"`
	Width     int    `json:"w,omitempty"`
	Height    int    `json:"h,omitempty"`
	Format    string `json:"fmt,omitempty"`
	Size      int    `json:"size,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`

	Data []byte `json:"-"`
}

type Video struct {
	Link   string `json:"link"`
	Format string `json:"fmt,omitempty"`
	Length int    `json:"len,omitempty"`
	Size   int    `json:"size,omitempty"`
	Image  *Image `json:"img,omitempty"`
	Name   string `json:"name,omitempty"`

	Data []byte `json:"-"`
}

type Audio struct {
	Link   string `json:"link"`
	Format string `json:"fmt,omitempty"`
	Length int    `json:"len,omitempty"`
	Size   int    `json:"size,omitempty"`
	Name   string `json:"name,omitempty"`

	Data []byte `json:"-"`
}

type File struct {
	Link   string `json:"link"`
	Name   string `json:"name"`
	Size   int    `json:"size,omitempty"`
	Format string `json:"fmt,omitempty"`

	Data []byte `json:"-"`
}

type WebPage struct {
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Image   *Image `json:"image,omitempty"`
	Link    string `json:"link"`
}
