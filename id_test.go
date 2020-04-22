package gox

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 256; i++ {
		id := NextID()
		t.Logf("%d %0X %s", id, id, id.ShortString())
		//time.Sleep(time.Millisecond * 1)
	}

	var id ID = 123
	if id.ShortString() != "1z" {
		t.Log(id.ShortString())
		t.FailNow()
	}

	id = 62
	if id.ShortString() != "10" {
		t.Log(id.ShortString())
		t.FailNow()
	}

	id = math.MaxInt64
	if i, _ := ParseShortID(id.ShortString()); i != id {
		t.Log(id.ShortString(), i)
		t.FailNow()
	}
}

func TestID_PrettyString(t *testing.T) {
	for i := 0; i < 256; i++ {
		id := NextID()
		t.Logf("%d %0X %s", id, id, id.PrettyString())
		//time.Sleep(time.Millisecond * 1)
	}

	var id ID = 123
	if id.PrettyString() != "4M" {
		t.Log(id.PrettyString())
		t.FailNow()
	}

	id = 34
	if id.PrettyString() != "21" {
		t.Log(id.PrettyString())
		t.FailNow()
	}

	id = math.MaxInt64
	if i, _ := ParsePrettyID(id.PrettyString()); i != id {
		t.Log(id.PrettyString(), i)
		t.FailNow()
	}
}

func TestNumberGetterFunc_GetNumber(t *testing.T) {
	ip := GetShardIDByIP()
	t.Logf("%0X", ip)
	i1 := KeepRightBits(ip, 8)
	t.Logf("%0X %d", i1, i1)

}
