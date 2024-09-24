package g

import (
	"github.com/telroshan/go-sfml/v2/graphics"
)

type MyCircle struct {
	x      float32
	y      float32
	radius float32

	red   byte
	green byte
	blue  byte
}

//------------------------------------------------------------//

func MyCircle_create() *MyCircle {
	o := MyCircle{
		x:      0,
		y:      0,
		radius: 50,
		red:    255,
		green:  255,
		blue:   255,
	}
	return &o
}
func MyCircle_destroy(o *MyCircle) {
}
func MyCircle_display(o *MyCircle, w *MyWindow) {
	circle := graphics.SfCircleShape_create()
	graphics.SfCircleShape_setRadius(circle, o.radius)
	graphics.SfCircleShape_setPosition(circle, MakeVector2(o.x, o.y))
	graphics.SfCircleShape_setFillColor(circle, graphics.SfColor_fromRGB(o.red, o.green, o.blue))
	graphics.SfRenderWindow_drawCircleShape(w.w, circle, GetNullRenderState())
	graphics.SfCircleShape_destroy(circle)
}

//------------------------------------------------------------//

func MyCircle_setX(o *MyCircle, x float32) {
	o.x = x
}
func MyCircle_setY(o *MyCircle, y float32) {
	o.y = y
}
func MyCircle_setRadius(o *MyCircle, r float32) {
	o.radius = r
}
func MyCircle_setColor(o *MyCircle, r, g, b byte) {
	o.red = r
	o.green = g
	o.blue = b
}
func MyCircle_setPosition(o *MyCircle, x, y float32) {
	o.x = x
	o.y = y
}

//------------------------------------------------------------//

func MyCircle_getRadius(o *MyCircle) float32 {
	return o.radius
}
func MyCircle_getColor(o *MyCircle) (byte, byte, byte) {
	return o.red, o.green, o.blue
}
func MyCircle_getX(o *MyCircle) float32 {
	return o.x
}
func MyCircle_getY(o *MyCircle) float32 {
	return o.y
}
