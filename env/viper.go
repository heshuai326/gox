package env

import (
	"time"

	"github.com/gopub/log"
	"github.com/spf13/viper"
)

type ViperManager struct {
}

func NewViperManager() *ViperManager {
	viper.AutomaticEnv()
	return &ViperManager{}
}

func (m *ViperManager) String(key string, defaultVal string) string {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetString(key)
	log.Debugf("Env: %s=%s", key, v)
	return v
}

func (m *ViperManager) Int(key string, defaultVal int) int {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetInt(key)
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func (m *ViperManager) Int64(key string, defaultVal int64) int64 {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetInt64(key)
	log.Debugf("Env: %s=%d", key, v)
	return v
}

func (m *ViperManager) Float64(key string, defaultVal float64) float64 {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetFloat64(key)
	log.Debugf("Env: %s=%f", key, v)
	return v
}

func (m *ViperManager) Duration(key string, defaultVal time.Duration) time.Duration {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetDuration(key)
	log.Debugf("Env: %s=%v", key, v)
	return v
}

func (m *ViperManager) Bool(key string, defaultVal bool) bool {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetBool(key)
	log.Debugf("Env: %s=%t", key, v)
	return v
}

func (m *ViperManager) IntSlice(key string, defaultVal []int) []int {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetIntSlice(key)
	log.Debugf("Env: %s=%v", key, v)
	return v
}

func (m *ViperManager) StringSlice(key string, defaultVal []string) []string {
	if viper.Get(key) == nil {
		log.Debugf("Env: missing %s", key)
	}
	viper.SetDefault(key, defaultVal)
	v := viper.GetStringSlice(key)
	log.Debugf("Env: %s=%v", key, v)
	return v
}
