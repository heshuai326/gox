package mobile

import (
	"github.com/gopub/gox"
	"github.com/gopub/log"
)

const (
	NoNetwork = 0
	Cellular  = 1
	WIFI      = 2
)

const (
	IDle       = 0
	Connecting = 1
	Connected  = 2
)

func GetOutboundIP() string {
	ip, err := gox.GetOutboundIP()
	if err != nil {
		log.Errorf("GetOutboundIP: %v", err)
		return ""
	}
	return ip.String()
}

func GetIPv4(interfaceName string) string {
	nameToIPList, err := gox.ListIPv4()
	if err != nil {
		log.Errorf("ListIPv4: %v", err)
		return ""
	}

	if l, ok := nameToIPList[interfaceName]; ok {
		return l[0]
	}
	return ""
}

func GetIPv6(interfaceName string) string {
	nameToIPList, err := gox.ListIPv6()
	if err != nil {
		log.Errorf("ListIPv4: %v", err)
		return ""
	}

	if l, ok := nameToIPList[interfaceName]; ok {
		return l[0]
	}
	return ""
}
