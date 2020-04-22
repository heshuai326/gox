package gox

type DeviceType string

func (d DeviceType) IsValid() bool {
	switch d {
	case IOS, Android, MacOS, Windows:
		return true
	default:
		return false
	}
}

const (
	IOS     DeviceType = "ios"
	Android DeviceType = "android"
	MacOS   DeviceType = "mac"
	Windows DeviceType = "windows"
)

type Env int

func (e Env) IsValid() bool {
	switch e {
	case Dev, Testing, Staging, Prod:
		return true
	default:
		return false
	}
}

func (e Env) String() string {
	switch e {
	case Dev:
		return "dev"
	case Testing:
		return "testing"
	case Staging:
		return "staging"
	case Prod:
		return "prod"
	default:
		return "nil"
	}
}

const (
	Dev Env = iota
	Testing
	Staging
	Prod
)
