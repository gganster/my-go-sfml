package libsfml

import (
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

type MyEvent struct {
	_ev window.SfEvent

	IsArrowUpPressed    int
	IsArrowDownPressed  int
	IsArrowLeftPressed  int
	IsArrowRightPressed int
	IsSpacePressed      int
}

func MyEvent_create() *MyEvent {
	o := MyEvent{window.NewSfEvent(), 0, 0, 0, 0, 0}
	return &o
}

func MyEvent_destroy(o *MyEvent) {
	window.DeleteSfEvent(o._ev)
}

func myEvent_reset(o *MyEvent) {
	o.IsArrowUpPressed = 0
	o.IsArrowDownPressed = 0
	o.IsArrowLeftPressed = 0
	o.IsArrowRightPressed = 0
	o.IsSpacePressed = 0
}

func MyEvent_update(o *MyEvent, w *graphics.Struct_SS_sfRenderWindow) {
	for window.SfWindow_pollEvent(*w, o._ev) > 0 {
		if o._ev.GetEvType() == window.SfEventType(window.SfEvtClosed) ||
			(o._ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) && o._ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyEscape)) {
			window.SfWindow_close(*w)
			break
		}

		switch o._ev.GetEvType() {
		case window.SfEventType(window.SfEvtKeyPressed):
			switch o._ev.GetKey().GetCode() {
			case window.SfKeyCode(window.SfKeyUp):
				o.IsArrowUpPressed = 1
			case window.SfKeyCode(window.SfKeyDown):
				o.IsArrowDownPressed = 1
			case window.SfKeyCode(window.SfKeyLeft):
				o.IsArrowLeftPressed = 1
			case window.SfKeyCode(window.SfKeyRight):
				o.IsArrowRightPressed = 1
			case window.SfKeyCode(window.SfKeySpace):
				o.IsSpacePressed = 1
			}

		case window.SfEventType(window.SfEvtKeyReleased):
			switch o._ev.GetKey().GetCode() {
			case window.SfKeyCode(window.SfKeyUp):
				o.IsArrowUpPressed = 0
			case window.SfKeyCode(window.SfKeyDown):
				o.IsArrowDownPressed = 0
			case window.SfKeyCode(window.SfKeyLeft):
				o.IsArrowLeftPressed = 0
			case window.SfKeyCode(window.SfKeyRight):
				o.IsArrowRightPressed = 0
			case window.SfKeyCode(window.SfKeySpace):
				o.IsSpacePressed = 0
			}
		}
	}
}

func MyEvent_dIsplayDebug(o *MyEvent) {
	println("IsArrowUpPressed: ", o.IsArrowUpPressed)
	println("IsArrowDownPressed: ", o.IsArrowDownPressed)
	println("IsArrowLeftPressed: ", o.IsArrowLeftPressed)
	println("IsArrowRightPressed: ", o.IsArrowRightPressed)
	println("IsSpacePressed: ", o.IsSpacePressed)
}
