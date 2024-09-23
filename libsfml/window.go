package libsfml

import (
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

type MyWindow struct {
	// -------- private --------//
	_vm window.SfVideoMode
	_cs window.SfContextSettings
	_w  graphics.Struct_SS_sfRenderWindow

	// -------- public --------//
	Ev *MyEvent
}

func MyWindow_create(title string, width uint, height uint) *MyWindow {
	_vm := window.NewSfVideoMode()
	_vm.SetWidth(width)
	_vm.SetHeight(height)
	_vm.SetBitsPerPixel(32)

	_cs := window.NewSfContextSettings()

	_w := graphics.SfRenderWindow_create(_vm, title, uint(window.SfTitlebar|window.SfClose), _cs)
	window.SfWindow_setVerticalSyncEnabled(_w, 1)

	Ev := MyEvent_create()

	o := MyWindow{_vm, _cs, _w, Ev}

	return &o
}

func MyWindow_destroy(o *MyWindow) {
	window.DeleteSfVideoMode(o._vm)
	window.DeleteSfContextSettings(o._cs)
	window.SfWindow_destroy(o._w)
	MyEvent_destroy(o.Ev)
}

func MyWindow_update(o *MyWindow) int {
	isOpen := window.SfWindow_isOpen(o._w)

	if isOpen == CFalse {
		return CFalse
	}

	MyEvent_update(o.Ev, &o._w)

	return CTrue
}

func MyWindow_clear(o *MyWindow) {
	graphics.SfRenderWindow_clear(o._w, graphics.SfColor_fromRGB(0, 0, 0))
}

func MyWindow_display(o *MyWindow) {
	graphics.SfRenderWindow_display(o._w)
}
