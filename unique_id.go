package gox

import (
	"fmt"
	"math/rand"
	osuser "os/user"
	"strings"
	"time"
)

func randomString() string {
	b := &strings.Builder{}
	addrs, err := GetMacAddrs()
	if err == nil {
		for _, a := range addrs {
			b.WriteString(a)
		}
	}

	u, err := osuser.Current()
	if err == nil {
		b.WriteString(u.Name)
		b.WriteString(u.Username)
		b.WriteString(u.Gid)
		b.WriteString(u.HomeDir)
		b.WriteString(u.Uid)
	}

	if ip, err := GetOutboundIP(); err == nil {
		b.WriteString(ip.String())
	}
	b.WriteString(time.Now().String())
	b.WriteString(NextID().ShortString())
	b.WriteString(fmt.Sprint(rand.Int()))
	return b.String()
}

// UniqueID returns an unique id
// Deprecated: please use UniqueID32/UniqueID40/UniqueID64 instead
func UniqueID() string {
	return SHA1(randomString())
}

// UniqueID32 returns an unique id of 32 letters
func UniqueID32() string {
	return MD5(randomString())
}

// UniqueID40 returns an unique id of 40 letters
func UniqueID40() string {
	return SHA1(randomString())
}

// UniqueID64 returns an unique id of 64 letters
func UniqueID64() string {
	return SHA256(randomString())
}
