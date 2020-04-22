package gox_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gopub/gox"
)

func nextImage() *gox.Image {
	return &gox.Image{
		Link:   "https://www.image.com/" + fmt.Sprint(time.Now().Unix()),
		Width:  200,
		Height: 800,
		Format: "png",
	}
}

func nextVideo() *gox.Video {
	return &gox.Video{
		Link:   "http://www.video.com/" + fmt.Sprint(time.Now().Unix()),
		Format: "rmvb",
		Length: 1230,
		Size:   90,
		Image:  nextImage(),
	}
}

func TestID(t *testing.T) {
	var v gox.ID = 10
	if err := gox.RegisterAny(gox.ID(0)); err != nil {
		t.Error(err)
		t.FailNow()
	}
	a := gox.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *gox.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Val().(gox.ID); !ok {
		t.Error("expected gox.ID")
		t.FailNow()
	} else if v != v2 {
		t.Error("expected equal gox.ID")
		t.FailNow()
	}
}

func TestText(t *testing.T) {
	v := "hello"
	a := gox.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *gox.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Val().(string); !ok {
		t.Error("expected Text")
		t.FailNow()
	} else if v != v2 {
		t.Error("expected equal text value")
		t.FailNow()
	}
}

func TestImage(t *testing.T) {
	v := nextImage()
	a := gox.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *gox.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Val().(*gox.Image); !ok {
		t.Error("expected Image")
		t.FailNow()
	} else if *v != *v2 {
		t.Error("expected equal image value")
		t.FailNow()
	}
}

func TestVideo(t *testing.T) {
	v := nextVideo()
	a := gox.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *gox.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Val().(*gox.Video); !ok {
		t.Error("expected Video")
		t.FailNow()
	} else if *v.Image != *v2.Image || v.Link != v2.Link || v.Size != v2.Size || v.Format != v2.Format || v.Length != v2.Length {
		t.Error("expected equal video value")
		t.FailNow()
	}
}

func TestArray(t *testing.T) {

	var items []*gox.Any
	items = append(items, gox.NewAny("hello"))
	items = append(items, gox.NewAny(nextImage()))
	items = append(items, gox.NewAny(nextVideo()))
	b, err := json.Marshal(items)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(b))

	var items2 []*gox.Any
	err = json.Unmarshal(b, &items2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if items[0].Val().(string) != items2[0].Val().(string) {
		t.FailNow()
	}

	//if *items[1].Val().(*gox.Image) != *items2[1].Val().(*gox.Image) {
	//	t.FailNow()
	//}
	//
	//{
	//	v := items[2].Val().(*gox.Video)
	//	v2 := items2[2].Val().(*gox.Video)
	//	if *v.Image != *v2.Image || v.Link != v2.Link || v.Size != v2.Size || v.Format != v2.Format || v.Length != v2.Length {
	//		t.Error("expected equal video value")
	//		t.FailNow()
	//	}
	//}
}
