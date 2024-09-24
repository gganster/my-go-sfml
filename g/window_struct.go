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

	// -------- for mouse --------//

	mouseX              int
	mouseY              int
	mouseLeftClick      int
	mouseRightClick     int
	mouseLeftClickOnce  int
	mouseRightClickOnce int

	// -------- for keyboard --------//

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

	// -------- for pressedOnce --------//
	pressedOnceArrowUp    int
	pressedOnceArrowDown  int
	pressedOnceArrowLeft  int
	pressedOnceArrowRight int
	pressedOnceSpace      int

	pressedOnceA int
	pressedOnceB int
	pressedOnceC int
	pressedOnceD int
	pressedOnceE int
	pressedOnceF int
	pressedOnceG int
	pressedOnceH int
	pressedOnceI int
	pressedOnceJ int
	pressedOnceK int
	pressedOnceL int
	pressedOnceM int
	pressedOnceN int
	pressedOnceO int
	pressedOnceP int
	pressedOnceQ int
	pressedOnceR int
	pressedOnceS int
	pressedOnceT int
	pressedOnceU int
	pressedOnceV int
	pressedOnceW int
	pressedOnceX int
	pressedOnceY int
	pressedOnceZ int
}
