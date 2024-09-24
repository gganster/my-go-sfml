package g

import (
	"github.com/telroshan/go-sfml/v2/graphics"
)

type MyRect struct {
	x      float32
	y      float32
	width  float32
	height float32

	red   byte
	green byte
	blue  byte
}

func MyRect_create() *MyRect {
	o := MyRect{
		x:      0,
		y:      0,
		width:  50,
		height: 50,
		red:    255,
		green:  255,
		blue:   255,
	}
	return &o
}

//------------------------------------------------------------//

func MyRect_destroy(o *MyRect) {
}

func MyRect_display(o *MyRect, w *MyWindow) {
	rect := graphics.SfRectangleShape_create()
	graphics.SfRectangleShape_setSize(rect, MakeVector2(o.width, o.height))
	graphics.SfRectangleShape_setPosition(rect, MakeVector2(o.x, o.y))
	graphics.SfRectangleShape_setFillColor(rect, graphics.SfColor_fromRGB(o.red, o.green, o.blue))
	graphics.SfRenderWindow_drawRectangleShape(w.w, rect, GetNullRenderState())
	graphics.SfRectangleShape_destroy(rect)
}

//------------------------------------------------------------//

func MyRect_setX(o *MyRect, x float32) {
	o.x = x
}
func MyRect_setY(o *MyRect, y float32) {
	o.y = y
}
func MyRect_setWidth(o *MyRect, w float32) {
	o.width = w
}
func MyRect_setHeight(o *MyRect, h float32) {
	o.height = h
}
func MyRect_setColor(o *MyRect, r, g, b byte) {
	o.red = r
	o.green = g
	o.blue = b
}
func MyRect_setPosition(o *MyRect, x, y float32) {
	o.x = x
	o.y = y
}

//------------------------------------------------------------//

func MyRect_getWidth(o *MyRect) float32 {
	return o.width
}
func MyRect_getHeight(o *MyRect) float32 {
	return o.height
}
func MyRect_getPosition(o *MyRect) (float32, float32) {
	return o.x, o.y
}
func MyRect_getColor(o *MyRect) (byte, byte, byte) {
	return o.red, o.green, o.blue
}
func MyRect_getX(o *MyRect) float32 {
	return o.x
}
func MyRect_getY(o *MyRect) float32 {
	return o.y
}
