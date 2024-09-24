package g

import (
	"github.com/telroshan/go-sfml/v2/graphics"
)

type MyText struct {
	t graphics.Struct_SS_sfText
	f graphics.Struct_SS_sfFont

	x    float32
	y    float32
	size uint
	text string

	red   byte
	green byte
	blue  byte
}

//------------------------------------------------------------//

func MyText_create() *MyText {
	f := graphics.SfFont_createFromFile("g/tuffy.ttf")
	if f == nil || f.Swigcptr() == CFalse {
		panic("Couldn't load tuffy.ttf")
	}
	t := graphics.SfText_create()
	graphics.SfText_setFont(t, f)

	o := MyText{
		t:     t,
		f:     f,
		x:     0,
		y:     0,
		size:  25,
		text:  "Hello, World!",
		red:   255,
		green: 255,
		blue:  255,
	}
	return &o
}

func MyText_destroy(o *MyText) {
	graphics.SfText_destroy(o.t)
	graphics.SfFont_destroy(o.f)
}

func MyText_display(o *MyText, w *MyWindow) {
	graphics.SfText_setString(o.t, o.text)
	graphics.SfText_setCharacterSize(o.t, o.size)
	graphics.SfText_setPosition(o.t, MakeVector2(o.x, o.y))
	graphics.SfText_setFillColor(o.t, graphics.SfColor_fromRGB(o.red, o.green, o.blue))
	graphics.SfRenderWindow_drawText(w.w, o.t, GetNullRenderState())
}

//------------------------------------------------------------//

func MyText_setX(o *MyText, x float32) {
	o.x = x
}

func MyText_setY(o *MyText, y float32) {
	o.y = y
}

func MyText_setSize(o *MyText, s uint) {
	o.size = s
}

func MyText_setText(o *MyText, t string) {
	o.text = t
}

func MyText_setColor(o *MyText, r, g, b byte) {
	o.red = r
	o.green = g
	o.blue = b
}

func MyText_setPosition(o *MyText, x, y float32) {
	o.x = x
	o.y = y
}

//------------------------------------------------------------//

func MyText_getX(o *MyText) float32 {
	return o.x
}

func MyText_getY(o *MyText) float32 {
	return o.y
}

func MyText_getSize(o *MyText) uint {
	return o.size
}

func MyText_getText(o *MyText) string {
	return o.text
}

func MyText_getColor(o *MyText) (byte, byte, byte) {
	return o.red, o.green, o.blue
}
