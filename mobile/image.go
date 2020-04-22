package mobile

import "github.com/gopub/gox"

type Image gox.Image

func NewImage() *Image {
	return new(Image)
}

type ImageList struct {
	List []*gox.Image
}

func (l *ImageList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *ImageList) Get(index int) *gox.Image {
	return l.List[index]
}
