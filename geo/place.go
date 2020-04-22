package geo

import (
	"database/sql/driver"
	"fmt"

	"github.com/gopub/gox/sql"
)

// Place
type Place struct {
	Code     string `json:"code,omitempty"`
	Name     string `json:"name,omitempty"`
	Location *Point `json:"point,omitempty"`
}

func NewPlace() *Place {
	return &Place{}
}

var _ driver.Valuer = (*Place)(nil)
var _ sql.Scanner = (*Place)(nil)

func (p *Place) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("cannot parse %v into string", src)
	}
	if s == "" {
		return nil
	}
	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}
	if len(fields) != 3 {
		return fmt.Errorf("parse composite fields %s", s)
	}
	p.Code = fields[0]
	p.Name = fields[1]
	if len(fields[2]) > 0 {
		p.Location = new(Point)
		if err := p.Location.Scan(fields[2]); err != nil {
			return fmt.Errorf("scan place.location: %w", err)
		}
	}
	return nil
}

func (p *Place) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	loc, err := p.Location.Value()
	if err != nil {
		return nil, fmt.Errorf("get location value: %w", err)
	}
	if locStr, ok := loc.(string); ok {
		loc = sql.Escape(locStr)
	}
	s := fmt.Sprintf("(%s,%s,%s)", sql.Escape(p.Code), sql.Escape(p.Name), loc)
	return s, nil
}
