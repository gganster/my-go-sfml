package g

import (
	"github.com/telroshan/go-sfml/v2/window"
)

func MyWindow_update(o *MyWindow) int {
	for window.SfWindow_pollEvent(o.w, o.ev) > 0 {
		if o.ev.GetEvType() == window.SfEventType(window.SfEvtClosed) ||
			(o.ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) && o.ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyEscape)) {
			window.SfWindow_close(o.w)
			return CFalse
		}

		switch o.ev.GetEvType() {
		case window.SfEventType(window.SfEvtKeyPressed):

			switch o.ev.GetKey().GetCode() {
			// Arrow + space
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
			// A-Z
			case window.SfKeyCode(window.SfKeyA):
				o.isAPressed = 1
			case window.SfKeyCode(window.SfKeyB):
				o.isBPressed = 1
			case window.SfKeyCode(window.SfKeyC):
				o.isCPressed = 1
			case window.SfKeyCode(window.SfKeyD):
				o.isDPressed = 1
			case window.SfKeyCode(window.SfKeyE):
				o.isEPressed = 1
			case window.SfKeyCode(window.SfKeyF):
				o.isFPressed = 1
			case window.SfKeyCode(window.SfKeyG):
				o.isGPressed = 1
			case window.SfKeyCode(window.SfKeyH):
				o.isHPressed = 1
			case window.SfKeyCode(window.SfKeyI):
				o.isIPressed = 1
			case window.SfKeyCode(window.SfKeyJ):
				o.isJPressed = 1
			case window.SfKeyCode(window.SfKeyK):
				o.isKPressed = 1
			case window.SfKeyCode(window.SfKeyL):
				o.isLPressed = 1
			case window.SfKeyCode(window.SfKeyM):
				o.isMPressed = 1
			case window.SfKeyCode(window.SfKeyN):
				o.isNPressed = 1
			case window.SfKeyCode(window.SfKeyO):
				o.isOPressed = 1
			case window.SfKeyCode(window.SfKeyP):
				o.isPPressed = 1
			case window.SfKeyCode(window.SfKeyQ):
				o.isQPressed = 1
			case window.SfKeyCode(window.SfKeyR):
				o.isRPressed = 1
			case window.SfKeyCode(window.SfKeyS):
				o.isSPressed = 1
			case window.SfKeyCode(window.SfKeyT):
				o.isTPressed = 1
			case window.SfKeyCode(window.SfKeyU):
				o.isUPressed = 1
			case window.SfKeyCode(window.SfKeyV):
				o.isVPressed = 1
			case window.SfKeyCode(window.SfKeyW):
				o.isWPressed = 1
			case window.SfKeyCode(window.SfKeyX):
				o.isXPressed = 1
			case window.SfKeyCode(window.SfKeyY):
				o.isYPressed = 1
			case window.SfKeyCode(window.SfKeyZ):
				o.isZPressed = 1
			}

		case window.SfEventType(window.SfEvtKeyReleased):
			switch o.ev.GetKey().GetCode() {
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
			// A-Z
			case window.SfKeyCode(window.SfKeyA):
				o.isAPressed = 0
			case window.SfKeyCode(window.SfKeyB):
				o.isBPressed = 0
			case window.SfKeyCode(window.SfKeyC):
				o.isCPressed = 0
			case window.SfKeyCode(window.SfKeyD):
				o.isDPressed = 0
			case window.SfKeyCode(window.SfKeyE):
				o.isEPressed = 0
			case window.SfKeyCode(window.SfKeyF):
				o.isFPressed = 0
			case window.SfKeyCode(window.SfKeyG):
				o.isGPressed = 0
			case window.SfKeyCode(window.SfKeyH):
				o.isHPressed = 0
			case window.SfKeyCode(window.SfKeyI):
				o.isIPressed = 0
			case window.SfKeyCode(window.SfKeyJ):
				o.isJPressed = 0
			case window.SfKeyCode(window.SfKeyK):
				o.isKPressed = 0
			case window.SfKeyCode(window.SfKeyL):
				o.isLPressed = 0
			case window.SfKeyCode(window.SfKeyM):
				o.isMPressed = 0
			case window.SfKeyCode(window.SfKeyN):
				o.isNPressed = 0
			case window.SfKeyCode(window.SfKeyO):
				o.isOPressed = 0
			case window.SfKeyCode(window.SfKeyP):
				o.isPPressed = 0
			case window.SfKeyCode(window.SfKeyQ):
				o.isQPressed = 0
			case window.SfKeyCode(window.SfKeyR):
				o.isRPressed = 0
			case window.SfKeyCode(window.SfKeyS):
				o.isSPressed = 0
			case window.SfKeyCode(window.SfKeyT):
				o.isTPressed = 0
			case window.SfKeyCode(window.SfKeyU):
				o.isUPressed = 0
			case window.SfKeyCode(window.SfKeyV):
				o.isVPressed = 0
			case window.SfKeyCode(window.SfKeyW):
				o.isWPressed = 0
			case window.SfKeyCode(window.SfKeyX):
				o.isXPressed = 0
			case window.SfKeyCode(window.SfKeyY):
				o.isYPressed = 0
			case window.SfKeyCode(window.SfKeyZ):
				o.isZPressed = 0
			}
		}
	}
	return CTrue
}
