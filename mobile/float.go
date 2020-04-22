package mobile

import "encoding/json"

type Float64List struct {
	List []float64
}

func NewFloatList() *Float64List {
	l := &Float64List{}
	return l
}

func (l *Float64List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *Float64List) Get(index int) float64 {
	return l.List[index]
}

func (l *Float64List) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *Float64List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}
