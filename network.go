package gox

import (
	"fmt"
	"net"
)

// GetOutboundIP preferred outbound ip of this machine
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	if err = conn.Close(); err != nil {
		return nil, err
	}
	return localAddr.IP, nil
}

func GetMacAddrs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	addrs := make([]string, len(ifaces))
	for i, ifa := range ifaces {
		addrs[i] = ifa.HardwareAddr.String()
	}

	return addrs, nil
}

func ListIPv4() (map[string][]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("get net interfaces: %w", err)
	}
	res := make(map[string][]string)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, fmt.Errorf("get addrs %s: %w", i.Name, err)
		}
		for _, addr := range addrs {
			//var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				if ip4 := v.IP.To4(); ip4 != nil {
					res[i.Name] = append(res[i.Name], ip4.String())
				}
			case *net.IPAddr:
				//	if ip4 := v.IP.To4(); ip4 != nil {
				//		res[i.Name] = append(res[i.Name], ip4.String())
				//	}
			}
		}
	}
	return res, nil
}

func ListIPv6() (map[string][]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("get net interfaces: %w", err)
	}
	res := make(map[string][]string)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, fmt.Errorf("get addrs %s: %w", i.Name, err)
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if ip6 := v.IP.To16(); ip6 != nil {
					res[i.Name] = append(res[i.Name], ip6.String())
				}
			case *net.IPAddr:
				//	if ip4 := v.IP.To4(); ip4 != nil {
				//		res[i.Name] = append(res[i.Name], ip4.String())
				//	}
			}
		}
	}
	return res, nil
}
