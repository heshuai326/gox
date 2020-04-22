package gox

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
)

func Diff(v1, v2 interface{}) string {
	return cmp.Diff(v1, v2)
}

func DiffJSON(v1, v2 interface{}) string {
	j1, _ := json.Marshal(v1)
	j2, _ := json.Marshal(v2)
	return cmp.Diff(string(j1), string(j2))
}
