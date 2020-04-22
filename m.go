package gox

import (
	"encoding/json"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gopub/log"

	"github.com/nyaruka/phonenumbers"
)

// M is a special map which provides convenient methods
type M map[string]interface{}

func (m M) Values(key string) []interface{} {
	value := m[key]
	if value == nil {
		return []interface{}{}
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		length := v.Len()
		var values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			values[i] = v.Index(i).Interface()
		}
		return values
	default:
		return []interface{}{value}
	}
}

func (m M) Value(key string) interface{} {
	value := m[key]
	if value == nil {
		return nil
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() > 0 {
			return v.Index(0).Interface()
		}
		return nil
	default:
		return value
	}
}

func (m M) Contains(key string) bool {
	return m.Value(key) != nil
}

func (m M) ContainsString(key string) bool {
	switch m.Value(key).(type) {
	case string:
		return true
	default:
		return false
	}
}

func (m M) String(key string) string {
	switch v := m.Value(key).(type) {
	case string:
		return v
	case json.Number:
		return string(v)
	default:
		return ""
	}
}

func (m M) TrimmedString(key string) string {
	switch v := m.Value(key).(type) {
	case string:
		return strings.TrimSpace(v)
	case json.Number:
		return string(v)
	default:
		return ""
	}
}

func (m M) DefaultString(key string, defaultValue string) string {
	switch v := m.Value(key).(type) {
	case string:
		return v
	default:
		return defaultValue
	}
}

func (m M) MustString(key string) string {
	switch v := m.Value(key).(type) {
	case string:
		return v
	default:
		panic("No string value for key:" + key)
	}
}

func (m M) Strings(key string) []string {
	_, found := m[key]
	if !found {
		return nil
	}

	values := m.Values(key)
	result := []string{}
	for _, v := range values {
		if str, ok := v.(string); ok {
			result = append(result, str)
		}
	}

	return result
}

func (m M) ContainsBool(key string) bool {
	_, e := ParseBool(m.Value(key))
	return e == nil
}

func (m M) Bool(key string) bool {
	v, _ := ParseBool(m.Value(key))
	return v
}

func (m M) DefaultBool(key string, defaultValue bool) bool {
	if v, err := ParseBool(m.Value(key)); err == nil {
		return v
	}
	return defaultValue
}

func (m M) MustBool(key string) bool {
	if v, err := ParseBool(m.Value(key)); err == nil {
		return v
	}
	panic("No bool value for key:" + key)
}

func (m M) Int(key string) int {
	v, _ := ParseInt(m.Value(key))
	return int(v)
}

func (m M) DefaultInt(key string, defaultVal int) int {
	if v, err := ParseInt(m.Value(key)); err == nil {
		return int(v)
	}
	return defaultVal
}

func (m M) MustInt(key string) int {
	if v, err := ParseInt(m.Value(key)); err == nil {
		return int(v)
	}
	panic("No int value for key:" + key)
}

func (m M) Ints(key string) []int {
	values := m.Values(key)
	result := make([]int, 0, len(values))
	for _, v := range values {
		i, e := ParseInt(v)
		if e == nil {
			result = append(result, int(i))
		}
	}

	return result
}

func (m M) ContainsInt64(key string) bool {
	_, err := ParseInt(m.Value(key))
	return err == nil
}

func (m M) Int64(key string) int64 {
	v, _ := ParseInt(m.Value(key))
	return v
}

func (m M) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := ParseInt(m.Value(key)); err == nil {
		return v
	}
	return defaultVal
}

func (m M) MustInt64(key string) int64 {
	if v, err := ParseInt(m.Value(key)); err == nil {
		return v
	}
	panic("No int64 value for key:" + key)
}

func (m M) Int64s(key string) []int64 {
	values := m.Values(key)
	result := []int64{}
	for _, v := range values {
		i, e := ParseInt(v)
		if e == nil {
			result = append(result, i)
		}
	}

	return result
}

func (m M) ContainsFloat64(key string) bool {
	_, err := ParseFloat(m.Value(key))
	return err == nil
}

func (m M) Float64(key string) float64 {
	v, _ := ParseFloat(m.Value(key))
	return v
}

func (m M) DefaultFloat64(key string, defaultValue float64) float64 {
	if v, err := ParseFloat(m.Value(key)); err == nil {
		return v
	}
	return defaultValue
}

func (m M) MustFloat64(key string) float64 {
	if v, err := ParseFloat(m.Value(key)); err == nil {
		return v
	}
	panic("No float64 value for key:" + key)
}

func (m M) Float64s(key string) []float64 {
	values := m.Values(key)
	result := []float64{}
	for _, val := range values {
		i, e := ParseFloat(val)
		if e == nil {
			result = append(result, i)
		}
	}

	return result
}

func (m M) BigInt(key string) *big.Int {
	s := m.String(key)
	n := big.NewInt(0)
	_, ok := n.SetString(s, 10)
	if !ok {
		_, ok = n.SetString(s, 16)
	}

	if ok {
		return n
	}

	if k, err := ParseInt(m[key]); err == nil {
		return big.NewInt(k)
	}

	return nil
}

func (m M) DefaultBigInt(key string, defaultVal *big.Int) *big.Int {
	if n := m.BigInt(key); n != nil {
		return n
	}
	return defaultVal
}

func (m M) MustBigInt(key string) *big.Int {
	if n := m.BigInt(key); n != nil {
		return n
	}
	panic("No big.Int value for key:" + key)
}

func (m M) BigFloat(key string) *big.Float {
	s := m.String(key)
	n := big.NewFloat(0)
	_, ok := n.SetString(s)
	if ok {
		return n
	}

	if k, err := ParseFloat(m[key]); err == nil {
		return big.NewFloat(k)
	}

	return nil
}

func (m M) DefaultBigFloat(key string, defaultVal *big.Float) *big.Float {
	if n := m.BigFloat(key); n != nil {
		return n
	}
	return defaultVal
}

func (m M) MustBigFloat(key string) *big.Float {
	if n := m.BigFloat(key); n != nil {
		return n
	}
	panic("No big.Float value for key:" + key)
}

func (m M) ContainsDecimal(key string) bool {
	s := m.String(key)
	if s == "" {
		return false
	}
	_, err := decimal.NewFromString(s)
	return err == nil
}

func (m M) Decimal(key string) decimal.Decimal {
	v, _ := decimal.NewFromString(m.String(key))
	return v
}

func (m M) DefaultDecimal(key string, defaultVal decimal.Decimal) decimal.Decimal {
	if v, err := decimal.NewFromString(m.String(key)); err == nil {
		return v
	}
	return defaultVal
}

func (m M) MustDecimal(key string) decimal.Decimal {
	if v, err := decimal.NewFromString(m.String(key)); err == nil {
		return v
	}
	panic("No decimal value for key:" + key)
}

func (m M) Map(key string) M {
	switch val := m.Value(key).(type) {
	case M:
		return val
	case map[string]interface{}:
		return M(val)
	default:
		return M{}
	}
}

func (m M) Date(key string) (time.Time, bool) {
	return m.DateInLocation(key, time.UTC)
}

func (m M) DateInLocation(key string, location *time.Location) (time.Time, bool) {
	str := strings.TrimSpace(m.String(key))
	if len(str) > 0 {
		birthday, err := time.ParseInLocation("2006-01-02", str, location)
		return birthday, err == nil
	}

	return time.Time{}, false
}

func (m M) PhoneNumber(key string) *PhoneNumber {
	switch v := m[key].(type) {
	case string:
		pn, err := phonenumbers.Parse(strings.TrimSpace(v), "")
		if err != nil || !phonenumbers.IsValidNumber(pn) {
			return nil
		}
		return &PhoneNumber{
			Code:      int(pn.GetCountryCode()),
			Number:    int64(pn.GetNationalNumber()),
			Extension: pn.GetExtension(),
		}
	case M, map[string]interface{}:
		data, err := json.Marshal(v)
		if err != nil {
			log.Errorf("Marshal failed: %v", err)
		}
		var pn *PhoneNumber
		err = json.Unmarshal(data, &pn)
		if err != nil {
			log.Errorf("Unmarshal failed: %v", err)
		}
		return pn
	default:
		return nil
	}
}

func (m M) Email(key string) string {
	s := m.String(key)
	s = strings.TrimSpace(s)
	if emailRegexp.MatchString(s) {
		return s
	}
	return ""
}

func (m M) URL(key string) string {
	s := m.String(key)
	s = strings.TrimSpace(s)
	_, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return s
}

func (m M) set(k string, v interface{}) {
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return
	}
	if val.IsZero() {
		if _, ok := m[k]; ok {
			return
		}
	}
	m[k] = v
}

func (m M) AddMap(val M) {
	for k, v := range val {
		m.set(k, v)
	}
}

func (m M) AddMapObj(obj interface{}) {
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		return
	}

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Map {
		return
	}

	length := len(v.MapKeys())
	if length == 0 {
		return
	}

	if v.MapKeys()[0].Kind() != reflect.String {
		panic("not map[string]interface{}")
	}

	for _, key := range v.MapKeys() {
		val := v.MapIndex(key).Interface()
		if val != nil {
			m.set(key.String(), val)
		}
	}
}

func (m M) JSON() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(data)
}

func (m M) RemoveEmptyValues() {
	m.RemoveSomeEmptyValues(nil)
}

func (m M) RemoveSomeEmptyValues(keys []string) {
	for k, v := range m {
		if len(keys) > 0 && indexOfStr(keys, k) < 0 {
			continue
		}
		val := reflect.ValueOf(v)
		rm := false
		switch val.Kind() {
		case reflect.Invalid:
			rm = true
		case reflect.String:
			rm = val.Len() == 0
		case reflect.Ptr:
			rm = val.IsNil()
		case reflect.Slice, reflect.Array, reflect.Chan:
			rm = val.IsNil() || val.Len() == 0
		case reflect.Map:
			if m1, ok1 := v.(map[string]interface{}); ok1 {
				M(m1).RemoveSomeEmptyValues(keys)
			} else if m2, ok2 := v.(M); ok2 {
				m2.RemoveEmptyValues()
			}

			if val.IsNil() {
				rm = true
			} else if val.Len() == 0 {
				rm = true
			}
		}

		if rm {
			delete(m, k)
		}
	}
}

func (m M) RemoveKeys(keys []string) {
	for k := range m {
		if indexOfStr(keys, k) < 0 {
			delete(m, k)
		}
	}
}

func (m M) RemoveAllExceptKeys(keys []string) {
	for k := range m {
		if indexOfStr(keys, k) < 0 {
			delete(m, k)
		}
	}
}

func indexOfStr(strs []string, s string) int {
	for i, str := range strs {
		if s == str {
			return i
		}
	}
	return -1
}
