package gox

import "encoding/json"

type Void struct{}

type Set struct {
	items map[interface{}]Void
}

func NewSet(capacity int) *Set {
	s := &Set{}
	s.items = make(map[interface{}]Void, capacity)
	return s
}

func (s *Set) Add(item interface{}) {
	s.items[item] = Void{}
}

func (s *Set) Contains(item interface{}) bool {
	_, found := s.items[item]
	return found
}

func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *Set) Slice() []interface{} {
	l := make([]interface{}, 0, len(s.items))
	for k := range s.items {
		l = append(l, k)
	}

	return l
}

type Int64Set struct {
	items map[int64]Void
}

var _ json.Unmarshaler = (*Int64Set)(nil)
var _ json.Marshaler = (*Int64Set)(nil)

func NewInt64Set(capacity int) *Int64Set {
	s := &Int64Set{}
	s.items = make(map[int64]Void, capacity)
	return s
}

func (s *Int64Set) Add(item int64) {
	s.items[item] = Void{}
}

func (s *Int64Set) Contains(item int64) bool {
	_, found := s.items[item]
	return found
}

func (s *Int64Set) Remove(item int64) {
	delete(s.items, item)
}

func (s *Int64Set) Slice() []int64 {
	l := make([]int64, 0, len(s.items))
	for k := range s.items {
		l = append(l, k)
	}

	return l
}

func (s *Int64Set) Map() map[int64]Void {
	return s.items
}

func (s *Int64Set) Size() int {
	return len(s.items)
}

func (s *Int64Set) UnmarshalJSON(data []byte) error {
	var ids []int64
	if err := json.Unmarshal(data, &ids); err != nil {
		return err
	}
	for _, id := range ids {
		s.Add(id)
	}
	return nil
}

func (s *Int64Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

type StringSet struct {
	items map[string]Void
}

var _ json.Unmarshaler = (*StringSet)(nil)
var _ json.Marshaler = (*StringSet)(nil)

func NewStringSet(capacity int) *StringSet {
	s := &StringSet{}
	s.items = make(map[string]Void, capacity)
	return s
}

func (s *StringSet) Add(item string) {
	s.items[item] = Void{}
}

func (s *StringSet) Contains(item string) bool {
	_, found := s.items[item]
	return found
}

func (s *StringSet) Remove(item string) {
	delete(s.items, item)
}

func (s *StringSet) Slice() []string {
	sl := make([]string, 0, len(s.items))
	for k := range s.items {
		sl = append(sl, k)
	}

	return sl
}

func (s *StringSet) Map() map[string]Void {
	return s.items
}

func (s *StringSet) Size() int {
	return len(s.items)
}

func (s *StringSet) UnmarshalJSON(data []byte) error {
	var l []string
	if err := json.Unmarshal(data, &l); err != nil {
		return err
	}
	for _, v := range l {
		s.Add(v)
	}
	return nil
}

func (s *StringSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}
