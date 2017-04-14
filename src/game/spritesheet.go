package gobject

import "github.com/gopherjs/gopherjs/js"

default_image_location = "imgs/default.png"

type SpriteSheet struct {

	Src string
	Delta float64
	Padding float64
	size float64
	count float64
}

func NewJS() *js.Object {
	return js.MakeWrapper(&New())
}

func New() SpriteSheet {
	return SpriteSheet{default_image_location,32,5,22,867}
}

func (s *SpriteSheet) SetSize(size float64) {
	if(size >= 0) {
		s.size = size
	}
}

func (s SpriteSheet) GetSize() float64 {
	return s.size
}

func (s *SpriteSheet) SetCount(count float64) {
	if (count >= 0) {
		s.count = count
	}
}

func (s SpriteSheet) GetCount() {
	return s.count
}
