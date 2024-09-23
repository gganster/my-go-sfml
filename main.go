package main

import (
	"sfml/libsfml"
)

func main() {
	w := libsfml.MyWindow_create("SFML Window", 800, 600)
	defer libsfml.MyWindow_destroy(w)

	r := libsfml.MyRect_create()
	defer libsfml.MyRect_destroy(r)

	for libsfml.MyWindow_update(w) == libsfml.CTrue {
		libsfml.MyWindow_clear(w)

		if (w.Ev.IsArrowUpPressed == 1) && (r.Y > 0) {
			r.Y -= 3
		}
		if (w.Ev.IsArrowDownPressed == 1) && (r.Y < 600-50) {
			r.Y += 3
		}
		if (w.Ev.IsArrowLeftPressed == 1) && (r.X > 0) {
			r.X -= 3
		}
		if (w.Ev.IsArrowRightPressed == 1) && (r.X < 800-50) {
			r.X += 3
		}

		libsfml.MyRect_display(r, w)
		libsfml.MyWindow_display(w)
	}
}
