package env

import (
	"time"
)

type Manager interface {
	String(key string, defaultVal string) string
	Int(key string, defaultVal int) int
	Int64(key string, defaultVal int64) int64
	Float64(key string, defaultVal float64) float64
	Duration(key string, defaultVal time.Duration) time.Duration
	Bool(key string, defaultVal bool) bool
	IntSlice(key string, defaultVal []int) []int
	StringSlice(key string, defaultVal []string) []string
}

var DefaultManager Manager = NewViperManager()

func String(key string, defaultVal string) string {
	return DefaultManager.String(key, defaultVal)
}

func Int(key string, defaultVal int) int {
	return DefaultManager.Int(key, defaultVal)
}

func Int64(key string, defaultVal int64) int64 {
	return DefaultManager.Int64(key, defaultVal)
}

func Float64(key string, defaultVal float64) float64 {
	return DefaultManager.Float64(key, defaultVal)
}

func Duration(key string, defaultVal time.Duration) time.Duration {
	return DefaultManager.Duration(key, defaultVal)
}

func Bool(key string, defaultVal bool) bool {
	return DefaultManager.Bool(key, defaultVal)
}

func IntSlice(key string, defaultVal []int) []int {
	return DefaultManager.IntSlice(key, defaultVal)
}

func StringSlice(key string, defaultVal []string) []string {
	return DefaultManager.StringSlice(key, defaultVal)
}
