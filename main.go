package main

import (
	. "sfml/g"
)

func main() {
	w := MyWindow_create("SFML Window", 800, 600)
	defer MyWindow_destroy(w)

	c := MyCircle_create()
	MyCircle_setPosition(c, 400, 300)
	MyCircle_setColor(c, 255, 0, 0)
	MyCircle_setRadius(c, 15)
	defer MyCircle_destroy(c)

	//boucle de rendu
	for MyWindow_update(w) == CTrue {

		//traite les évènements et met à jour l'état
		if MyWindow_isMouseLeftPressedOnce(w) {
			MyCircle_setPosition(
				c,
				MyWindow_getMouseX(w),
				MyWindow_getMouseY(w),
			)
		}

		//effacement de la fenêtre
		MyWindow_clear(w)

		//dessine les éléments à l'écran
		MyCircle_display(c, w)

		//affiche le rendu
		MyWindow_display(w)
	}
}
