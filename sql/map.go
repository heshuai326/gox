package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Map map[string]interface{}

var _ sql.Scanner = (*Map)(nil)
var _ driver.Valuer = Map(nil)

func (m *Map) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		s, ok := src.(string)
		if ok {
			b = []byte(s)
		}
	}

	if !ok {
		return fmt.Errorf("parse %v into sql.Map failed", src)
	}

	return json.Unmarshal(b, m)
}

func (m Map) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}
