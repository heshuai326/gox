package geo

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/golang/geo/s2"
	"github.com/gopub/gox/sql"
)

const (
	PI          = 3.14159265
	EarthRadius = 6_378_100
	EarthCircle = 2 * PI * EarthRadius
	Degree      = EarthCircle * 1000 / 360
)

type Point struct {
	X float64 `json:"x"` // X is longitude for geodetic coordinate
	Y float64 `json:"y"` // Y is latitude for geodetic coordinate
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

var _ driver.Valuer = (*Point)(nil)
var _ sql.Scanner = (*Point)(nil)

func (p *Point) Scan(src interface{}) error {
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
	if len(fields) == 1 {
		fields = strings.Split(fields[0], " ")
	}
	if len(fields) != 2 {
		return fmt.Errorf("parse composite fields %s", s)
	}
	_, err = fmt.Sscanf(fields[0], "%f", &p.X)
	if err != nil {
		return fmt.Errorf("parse x %s: %w", fields[0], err)
	}
	_, err = fmt.Sscanf(fields[1], "%f", &p.Y)
	if err != nil {
		return fmt.Errorf("parse y %s: %w", fields[1], err)
	}
	return nil
}

func (p *Point) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	v := fmt.Sprintf("POINT(%f %f)", p.X, p.Y)
	return v, nil
}

func (p *Point) Distance(v *Point) int {
	p1 := s2.PointFromLatLng(s2.LatLngFromDegrees(p.Y, p.X))
	p2 := s2.PointFromLatLng(s2.LatLngFromDegrees(v.Y, v.X))
	d := p1.Distance(p2)
	return int(d.Radians() * 6371000)
}
