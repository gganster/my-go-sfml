package g

import (
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

type MyWindow struct {
	// -------- private --------//
	vm window.SfVideoMode
	cs window.SfContextSettings
	w  graphics.Struct_SS_sfRenderWindow
	ev window.SfEvent

	IsArrowUpPressed    int
	IsArrowDownPressed  int
	IsArrowLeftPressed  int
	IsArrowRightPressed int
	IsSpacePressed      int

	isAPressed int
	isBPressed int
	isCPressed int
	isDPressed int
	isEPressed int
	isFPressed int
	isGPressed int
	isHPressed int
	isIPressed int
	isJPressed int
	isKPressed int
	isLPressed int
	isMPressed int
	isNPressed int
	isOPressed int
	isPPressed int
	isQPressed int
	isRPressed int
	isSPressed int
	isTPressed int
	isUPressed int
	isVPressed int
	isWPressed int
	isXPressed int
	isYPressed int
	isZPressed int
}

func MyWindow_create(title string, width uint, height uint) *MyWindow {
	vm := window.NewSfVideoMode()
	vm.SetWidth(width)
	vm.SetHeight(height)
	vm.SetBitsPerPixel(32)

	cs := window.NewSfContextSettings()

	w := graphics.SfRenderWindow_create(vm, title, uint(window.SfTitlebar|window.SfClose), cs)
	window.SfWindow_setVerticalSyncEnabled(w, 1)

	ev := window.NewSfEvent()

	o := MyWindow{
		vm,
		cs,
		w,
		ev,
		0, 0, 0, 0, 0, //arrows and space
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //letters
	}

	return &o
}

func MyWindow_destroy(o *MyWindow) {
	window.DeleteSfVideoMode(o.vm)
	window.DeleteSfContextSettings(o.cs)
	window.SfWindow_destroy(o.w)
	window.DeleteSfEvent(o.ev)
}

func MyWindow_clear(o *MyWindow) {
	graphics.SfRenderWindow_clear(o.w, graphics.SfColor_fromRGB(0, 0, 0))
}

func MyWindow_display(o *MyWindow) {
	graphics.SfRenderWindow_display(o.w)
}
