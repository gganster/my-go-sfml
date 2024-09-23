package libsfml

import (
	"github.com/telroshan/go-sfml/v2/graphics"
)

type MyRect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	red   byte
	green byte
	blue  byte
}

func MyRect_create() *MyRect {
	o := MyRect{
		X:      0,
		Y:      0,
		Width:  50,
		Height: 50,
		red:    255,
		green:  255,
		blue:   255,
	}
	return &o
}

func MyRect_destroy(o *MyRect) {
}

func MyRect_display(o *MyRect, w *MyWindow) {
	rect := graphics.SfRectangleShape_create()
	graphics.SfRectangleShape_setSize(rect, MakeVector2(o.Width, o.Height))
	graphics.SfRectangleShape_setPosition(rect, MakeVector2(o.X, o.Y))
	graphics.SfRectangleShape_setFillColor(rect, graphics.SfColor_fromRGB(o.red, o.green, o.blue))
	graphics.SfRenderWindow_drawRectangleShape(w._w, rect, GetNullRenderState())
	graphics.SfRectangleShape_destroy(rect)
}
