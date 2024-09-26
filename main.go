package main

import (
	"runtime"
	. "sfml/g"
	"time"
)

func init() { runtime.LockOSThread() }

func main() {
	w := MyWindow_create("SFML Window", 800, 600)
	defer MyWindow_destroy(w)

	t := MyText_create()
	defer MyText_destroy(t)
	MyText_setText(t, "Hello, World!")
	MyText_setSize(t, 50)
	MyText_setColor(t, 255, 0, 0)
	MyText_setPosition(t, 100, 100)

	//boucle de rendu
	nsec := time.Now().UnixNano()
	for MyWindow_update(w, &nsec) == CTrue {

		//effacement de la fenêtre
		MyWindow_clear(w)

		//dessine les éléments à l'écran
		MyText_display(t, w)

		//affiche le rendu
		MyWindow_display(w)
	}
}
