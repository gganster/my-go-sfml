package main

import (
	"fmt"
	. "sfml/g"
)

func main() {
	w := MyWindow_create("SFML Window", 800, 600)
	defer MyWindow_destroy(w)

	r := MyRect_create()
	defer MyRect_destroy(r)

	c := MyCircle_create()
	var cOrientationX float32 = 0.4
	var cOrientationY float32 = 0.4
	MyCircle_setPosition(c, 400, 300)
	MyCircle_setColor(c, 255, 0, 0)
	MyCircle_setRadius(c, 15)

	defer MyCircle_destroy(c)

	for MyWindow_update(w) == CTrue {
		MyWindow_clear(w)
		/*
			if (w.Ev.IsArrowUpPressed == 1) && (MyRect_getHeight(r) > 0) {
				MyRect_setHeight(r, MyRect_getHeight(r)+3)
			}
			if (w.Ev.IsArrowDownPressed == 1) && (MyRect_getHeight(r) < 600-50) {
				MyRect_setHeight(r, MyRect_getHeight(r)-3)
			}
			if (w.Ev.IsArrowLeftPressed == 1) && (MyRect_getWidth(r) > 0) {
				MyRect_setWidth(r, MyRect_getWidth(r)-3)
			}
			if (w.Ev.IsArrowRightPressed == 1) && (MyRect_getWidth(r) < 800-50) {
				MyRect_setWidth(r, MyRect_getWidth(r)+3)
			}
		*/
		if MyWindow_isArrowUpPressedOnce(w) {
			fmt.Println("Arrow Up pressed once")
		}
		//update circle position
		MyCircle_setPosition(
			c,
			MyCircle_getX(c)+cOrientationX*5,
			MyCircle_getY(c)+cOrientationY*5,
		)

		//bounce circle off walls
		if MyCircle_getX(c) < 0 || MyCircle_getX(c) > (800-MyCircle_getRadius(c)*2) {
			cOrientationX = -cOrientationX
		}
		if MyCircle_getY(c) < 0 || MyCircle_getY(c) > (600-MyCircle_getRadius(c)*2) {
			cOrientationY = -cOrientationY
		}

		MyCircle_display(c, w)
		MyRect_display(r, w)
		MyWindow_display(w)
	}
}
