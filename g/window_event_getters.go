package g

func MyWindow_getMouseX(o *MyWindow) float32        { return float32(o.mouseX) }
func MyWindow_getMouseY(o *MyWindow) float32        { return float32(o.mouseY) }
func MyWindow_isMouseLeftPressed(o *MyWindow) bool  { return (o.mouseLeftClick == 1) }
func MyWindow_isMouseRightPressed(o *MyWindow) bool { return (o.mouseRightClick == 1) }
func MyWindow_isMouseLeftPressedOnce(o *MyWindow) bool {
	if MyWindow_isMouseLeftPressed(o) && o.mouseLeftClickOnce == 0 {
		o.mouseLeftClickOnce = 1
		return true
	}
	return false
}
func MyWindow_isMouseRightPressedOnce(o *MyWindow) bool {
	if MyWindow_isMouseRightPressed(o) && o.mouseRightClickOnce == 0 {
		o.mouseRightClickOnce = 1
		return true
	}
	return false
}

func MyWindow_isArrowUpPressed(o *MyWindow) bool    { return (o.IsArrowUpPressed == 1) }
func MyWindow_isArrowDownPressed(o *MyWindow) bool  { return (o.IsArrowDownPressed == 1) }
func MyWindow_isArrowLeftPressed(o *MyWindow) bool  { return (o.IsArrowLeftPressed == 1) }
func MyWindow_isArrowRightPressed(o *MyWindow) bool { return (o.IsArrowRightPressed == 1) }
func MyWindow_isSpacePressed(o *MyWindow) bool      { return (o.IsSpacePressed == 1) }
func MyWindow_isAPressed(o *MyWindow) bool          { return (o.isAPressed == 1) }
func MyWindow_isBPressed(o *MyWindow) bool          { return (o.isBPressed == 1) }
func MyWindow_isCPressed(o *MyWindow) bool          { return (o.isCPressed == 1) }
func MyWindow_isDPressed(o *MyWindow) bool          { return (o.isDPressed == 1) }
func MyWindow_isEPressed(o *MyWindow) bool          { return (o.isEPressed == 1) }
func MyWindow_isFPressed(o *MyWindow) bool          { return (o.isFPressed == 1) }
func MyWindow_isGPressed(o *MyWindow) bool          { return (o.isGPressed == 1) }
func MyWindow_isHPressed(o *MyWindow) bool          { return (o.isHPressed == 1) }
func MyWindow_isIPressed(o *MyWindow) bool          { return (o.isIPressed == 1) }
func MyWindow_isJPressed(o *MyWindow) bool          { return (o.isJPressed == 1) }
func MyWindow_isKPressed(o *MyWindow) bool          { return (o.isKPressed == 1) }
func MyWindow_isLPressed(o *MyWindow) bool          { return (o.isLPressed == 1) }
func MyWindow_isMPressed(o *MyWindow) bool          { return (o.isMPressed == 1) }
func MyWindow_isNPressed(o *MyWindow) bool          { return (o.isNPressed == 1) }
func MyWindow_isOPressed(o *MyWindow) bool          { return (o.isOPressed == 1) }
func MyWindow_isPPressed(o *MyWindow) bool          { return (o.isPPressed == 1) }
func MyWindow_isQPressed(o *MyWindow) bool          { return (o.isQPressed == 1) }
func MyWindow_isRPressed(o *MyWindow) bool          { return (o.isRPressed == 1) }
func MyWindow_isSPressed(o *MyWindow) bool          { return (o.isSPressed == 1) }
func MyWindow_isTPressed(o *MyWindow) bool          { return (o.isTPressed == 1) }
func MyWindow_isUPressed(o *MyWindow) bool          { return (o.isUPressed == 1) }
func MyWindow_isVPressed(o *MyWindow) bool          { return (o.isVPressed == 1) }
func MyWindow_isWPressed(o *MyWindow) bool          { return (o.isWPressed == 1) }
func MyWindow_isXPressed(o *MyWindow) bool          { return (o.isXPressed == 1) }
func MyWindow_isYPressed(o *MyWindow) bool          { return (o.isYPressed == 1) }
func MyWindow_isZPressed(o *MyWindow) bool          { return (o.isZPressed == 1) }

func MyWindow_isArrowUpPressedOnce(o *MyWindow) bool {
	if MyWindow_isArrowUpPressed(o) && o.pressedOnceArrowUp == 0 {
		o.pressedOnceArrowUp = 1
		return true
	}
	return false
}
func MyWindow_isArrowDownPressedOnce(o *MyWindow) bool {
	if MyWindow_isArrowDownPressed(o) && o.pressedOnceArrowDown == 0 {
		o.pressedOnceArrowDown = 1
		return true
	}
	return false
}
func MyWindow_isArrowLeftPressedOnce(o *MyWindow) bool {
	if MyWindow_isArrowLeftPressed(o) && o.pressedOnceArrowLeft == 0 {
		o.pressedOnceArrowLeft = 1
		return true
	}
	return false
}
func MyWindow_isArrowRightPressedOnce(o *MyWindow) bool {
	if MyWindow_isArrowRightPressed(o) && o.pressedOnceArrowRight == 0 {
		o.pressedOnceArrowRight = 1
		return true
	}
	return false
}
func MyWindow_isSpacePressedOnce(o *MyWindow) bool {
	if MyWindow_isSpacePressed(o) && o.pressedOnceSpace == 0 {
		o.pressedOnceSpace = 1
		return true
	}
	return false
}
func MyWindow_isAPressedOnce(o *MyWindow) bool {
	if MyWindow_isAPressed(o) && o.pressedOnceA == 0 {
		o.pressedOnceA = 1
		return true
	}
	return false
}
func MyWindow_isBPressedOnce(o *MyWindow) bool {
	if MyWindow_isBPressed(o) && o.pressedOnceB == 0 {
		o.pressedOnceB = 1
		return true
	}
	return false
}
func MyWindow_isCPressedOnce(o *MyWindow) bool {
	if MyWindow_isCPressed(o) && o.pressedOnceC == 0 {
		o.pressedOnceC = 1
		return true
	}
	return false
}
func MyWindow_isDPressedOnce(o *MyWindow) bool {
	if MyWindow_isDPressed(o) && o.pressedOnceD == 0 {
		o.pressedOnceD = 1
		return true
	}
	return false
}
func MyWindow_isEPressedOnce(o *MyWindow) bool {
	if MyWindow_isEPressed(o) && o.pressedOnceE == 0 {
		o.pressedOnceE = 1
		return true
	}
	return false
}
func MyWindow_isFPressedOnce(o *MyWindow) bool {
	if MyWindow_isFPressed(o) && o.pressedOnceF == 0 {
		o.pressedOnceF = 1
		return true
	}
	return false
}
func MyWindow_isGPressedOnce(o *MyWindow) bool {
	if MyWindow_isGPressed(o) && o.pressedOnceG == 0 {
		o.pressedOnceG = 1
		return true
	}
	return false
}
func MyWindow_isHPressedOnce(o *MyWindow) bool {
	if MyWindow_isHPressed(o) && o.pressedOnceH == 0 {
		o.pressedOnceH = 1
		return true
	}
	return false
}
func MyWindow_isIPressedOnce(o *MyWindow) bool {
	if MyWindow_isIPressed(o) && o.pressedOnceI == 0 {
		o.pressedOnceI = 1
		return true
	}
	return false
}
func MyWindow_isJPressedOnce(o *MyWindow) bool {
	if MyWindow_isJPressed(o) && o.pressedOnceJ == 0 {
		o.pressedOnceJ = 1
		return true
	}
	return false
}
func MyWindow_isKPressedOnce(o *MyWindow) bool {
	if MyWindow_isKPressed(o) && o.pressedOnceK == 0 {
		o.pressedOnceK = 1
		return true
	}
	return false
}
func MyWindow_isLPressedOnce(o *MyWindow) bool {
	if MyWindow_isLPressed(o) && o.pressedOnceL == 0 {
		o.pressedOnceL = 1
		return true
	}
	return false
}
func MyWindow_isMPressedOnce(o *MyWindow) bool {
	if MyWindow_isMPressed(o) && o.pressedOnceM == 0 {
		o.pressedOnceM = 1
		return true
	}
	return false
}
func MyWindow_isNPressedOnce(o *MyWindow) bool {
	if MyWindow_isNPressed(o) && o.pressedOnceN == 0 {
		o.pressedOnceN = 1
		return true
	}
	return false
}
func MyWindow_isOPressedOnce(o *MyWindow) bool {
	if MyWindow_isOPressed(o) && o.pressedOnceO == 0 {
		o.pressedOnceO = 1
		return true
	}
	return false
}
func MyWindow_isPPressedOnce(o *MyWindow) bool {
	if MyWindow_isPPressed(o) && o.pressedOnceP == 0 {
		o.pressedOnceP = 1
		return true
	}
	return false
}
func MyWindow_isQPressedOnce(o *MyWindow) bool {
	if MyWindow_isQPressed(o) && o.pressedOnceQ == 0 {
		o.pressedOnceQ = 1
		return true
	}
	return false
}
func MyWindow_isRPressedOnce(o *MyWindow) bool {
	if MyWindow_isRPressed(o) && o.pressedOnceR == 0 {
		o.pressedOnceR = 1
		return true
	}
	return false
}
func MyWindow_isSPressedOnce(o *MyWindow) bool {
	if MyWindow_isSPressed(o) && o.pressedOnceS == 0 {
		o.pressedOnceS = 1
		return true
	}
	return false
}
func MyWindow_isTPressedOnce(o *MyWindow) bool {
	if MyWindow_isTPressed(o) && o.pressedOnceT == 0 {
		o.pressedOnceT = 1
		return true
	}
	return false
}
func MyWindow_isUPressedOnce(o *MyWindow) bool {
	if MyWindow_isUPressed(o) && o.pressedOnceU == 0 {
		o.pressedOnceU = 1
		return true
	}
	return false
}
func MyWindow_isVPressedOnce(o *MyWindow) bool {
	if MyWindow_isVPressed(o) && o.pressedOnceV == 0 {
		o.pressedOnceV = 1
		return true
	}
	return false
}
func MyWindow_isWPressedOnce(o *MyWindow) bool {
	if MyWindow_isWPressed(o) && o.pressedOnceW == 0 {
		o.pressedOnceW = 1
		return true
	}
	return false
}
func MyWindow_isXPressedOnce(o *MyWindow) bool {
	if MyWindow_isXPressed(o) && o.pressedOnceX == 0 {
		o.pressedOnceX = 1
		return true
	}
	return false
}
func MyWindow_isYPressedOnce(o *MyWindow) bool {
	if MyWindow_isYPressed(o) && o.pressedOnceY == 0 {
		o.pressedOnceY = 1
		return true
	}
	return false
}
func MyWindow_isZPressedOnce(o *MyWindow) bool {
	if MyWindow_isZPressed(o) && o.pressedOnceZ == 0 {
		o.pressedOnceZ = 1
		return true
	}
	return false
}
