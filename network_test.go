package gox

import "testing"

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(ip.String())
}
