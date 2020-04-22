package mobile

import (
	"encoding/json"

	"github.com/gopub/gox"
)

type Video gox.Video

func NewVideo() *Video {
	return new(Video)
}

func (v *Video) GetImage() *Image {
	return (*Image)(v.Image)
}

type Audio gox.Audio

func NewAudio() *Audio {
	return new(Audio)
}

type File gox.File

func NewFile() *File {
	return new(File)
}

type WebPage gox.WebPage

func NewWebPage() *WebPage {
	return new(WebPage)
}

type Any gox.Any

func (a *Any) TypeName() string {
	return (*gox.Any)(a).TypeName()
}

func (a *Any) SetImage(i *Image) {
	(*gox.Any)(a).SetImage((*gox.Image)(i))
}

func (a *Any) SetAudio(au *Audio) {
	(*gox.Any)(a).SetAudio((*gox.Audio)(au))
}

func (a *Any) SetVideo(v *Video) {
	(*gox.Any)(a).SetVideo((*gox.Video)(v))
}

func (a *Any) SetFile(f *File) {
	(*gox.Any)(a).SetFile((*gox.File)(f))
}

func (a *Any) SetWebPage(wp *WebPage) {
	(*gox.Any)(a).SetWebPage((*gox.WebPage)(wp))
}

func (a *Any) Image() *Image {
	return (*Image)((*gox.Any)(a).Image())
}

func (a *Any) Video() *Video {
	return (*Video)((*gox.Any)(a).Video())
}

func (a *Any) Audio() *Audio {
	return (*Audio)((*gox.Any)(a).Audio())
}

func (a *Any) File() *File {
	return (*File)((*gox.Any)(a).File())
}

func (a *Any) WebPage() *WebPage {
	return (*WebPage)((*gox.Any)(a).WebPage())
}

func NewAnyObj() *Any {
	return new(Any)
}

type AnyList struct {
	List []*gox.Any
}

func NewAnyListObj() *AnyList {
	return new(AnyList)
}

func NewAnyList(list []*gox.Any) *AnyList {
	return &AnyList{List: list}
}

func (a *AnyList) Size() int {
	if a == nil {
		return 0
	}
	return len(a.List)
}

func (a *AnyList) Get(index int) *Any {
	if a == nil {
		return nil
	}
	return (*Any)(a.List[index])
}

func (a *AnyList) Append(v *Any) {
	a.List = append(a.List, (*gox.Any)(v))
}

func (a *AnyList) Prepend(v *Any) {
	a.List = append([]*gox.Any{(*gox.Any)(v)}, a.List...)
}

func (a *AnyList) Insert(i int, v *Any) {
	if len(a.List) <= i {
		a.List = append(a.List, (*gox.Any)(v))
	} else {
		l := a.List[i:]
		l = append([]*gox.Any{(*gox.Any)(v)}, l...)
		a.List = append(a.List[0:i], l...)
	}
}

func (a *AnyList) RemoveAt(index int) {
	a.List = append(a.List[0:index], a.List[index+1:]...)
}

func (a *AnyList) Remove(v *Any) {
	i := a.IndexOf(v)
	if i >= 0 {
		a.RemoveAt(i)
	}
}

func (a *AnyList) IndexOf(v *Any) int {
	for i, m := range a.List {
		if (*Any)(m) == v {
			return i
		}
	}
	return -1
}

func (a *AnyList) FirstImage() *Image {
	for _, m := range a.List {
		if img := m.Image(); img != nil {
			return (*Image)(img)
		}
	}
	return nil
}

func (a *AnyList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &a.List)
}

func (a *AnyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.List)
}
