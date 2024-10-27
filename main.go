package main

import (
	"runtime"
	. "sfml/g"
	"time"
)

// constants
var RACKET_WIDTH float32 = 100
var RACKET_HEIGHT float32 = 10
var BALL_RADIUS float32 = 10

var wall [6][10]bool = [6][10]bool{
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
}

func abs(value float32) float32 {
	if value < 0 {
		return -value
	}
	return value
}

func isRectCollidingCircle(rect *MyRect, circle *MyCircle) bool {
	rx := MyRect_getX(rect)
	ry := MyRect_getY(rect)
	rw := MyRect_getWidth(rect)
	rh := MyRect_getHeight(rect)

	cx := MyCircle_getX(circle)
	cy := MyCircle_getY(circle)
	cw := MyCircle_getRadius(circle) * 2
	ch := MyCircle_getRadius(circle) * 2

	return rx < cx+cw &&
		rx+rw > cx &&
		ry < cy+ch &&
		ry+rh > cy
}

func main() {
	w := MyWindow_create("SFML Window", 800, 600)
	defer MyWindow_destroy(w)

	racket := MyRect_create()
	defer MyRect_destroy(racket)
	MyRect_setWidth(racket, RACKET_WIDTH)
	MyRect_setHeight(racket, RACKET_HEIGHT)
	MyRect_setColor(racket, 255, 0, 0)
	MyRect_setPosition(racket, 100, 550)

	ball := MyCircle_create()
	ballXSpeed := float32(3)
	ballYSpeed := float32(3)
	defer MyCircle_destroy(ball)
	MyCircle_setRadius(ball, BALL_RADIUS)
	MyCircle_setColor(ball, 255, 255, 255)
	MyCircle_setPosition(ball, 400, 300)

	brick := MyRect_create()
	defer MyRect_destroy(brick)
	MyRect_setWidth(brick, 100)
	MyRect_setHeight(brick, 50)
	MyRect_setColor(brick, 0, 255, 0)
	MyRect_setPosition(brick, 100, 100)

	nsec := time.Now().UnixNano()
	for MyWindow_update(w, &nsec) == CTrue {

		//handle brick collision
		for i := 0; i < 6; i++ {
			for j := 0; j < 10; j++ {
				if wall[i][j] {
					MyRect_setPosition(brick, float32(j*100), float32(i*50))
					ballX := MyCircle_getX(ball)
					ballY := MyCircle_getY(ball)
					radius := MyCircle_getRadius(ball)
					rectX := MyRect_getX(brick)
					rectY := MyRect_getY(brick)
					rectWidth := MyRect_getWidth(brick)
					rectHeight := MyRect_getHeight(brick)
					ballRight := ballX + radius*2
					ballTop := ballY
					ballBottom := ballY + radius*2
					ballLeft := ballX

					if ballRight >= rectX && ballLeft <= rectX+rectWidth &&
						ballBottom >= rectY && ballTop <= rectY+rectHeight {
						// goback slightly to avoid glitch
						MyCircle_setX(ball, MyCircle_getX(ball)-ballXSpeed)
						MyCircle_setY(ball, MyCircle_getY(ball)-ballYSpeed)

						// Déterminer si la collision est horizontale ou verticale
						if abs(ballX-(rectX+rectWidth/2)) > abs(ballY-(rectY+rectHeight/2)) {
							ballXSpeed = -ballXSpeed
							wall[i][j] = false
						} else {
							ballYSpeed = -ballYSpeed
							wall[i][j] = false
						}
					}
				}
			}
		}

		//handle racket movement
		if MyWindow_isArrowLeftPressed(w) && MyRect_getX(racket) > 0 {
			MyRect_setX(racket, MyRect_getX(racket)-4)
		}
		if MyWindow_isArrowRightPressed(w) && MyRect_getX(racket) < 800-RACKET_WIDTH {
			MyRect_setX(racket, MyRect_getX(racket)+4)
		}

		//handle ball movement
		if MyCircle_getX(ball) < 0 || MyCircle_getX(ball) > 800-BALL_RADIUS*2 { //bounce on top
			//reviens légèrement en arrière pour éviter de glitch
			MyCircle_setX(ball, MyCircle_getX(ball)-ballXSpeed)
			ballXSpeed = -ballXSpeed
		}
		if MyCircle_getY(ball) < 0 { //bounce on the sides
			//reviens légèrement en arrière pour éviter de glitch
			MyCircle_setY(ball, MyCircle_getY(ball)-ballYSpeed)
			ballYSpeed = -ballYSpeed
		}
		if MyCircle_getY(ball) > 600-BALL_RADIUS*2 { //perdu
			break
		}
		if isRectCollidingCircle(racket, ball) {
			//reviens légèrement en arrière pour éviter de glitch
			MyCircle_setY(ball, MyCircle_getY(ball)-ballYSpeed)
			ballYSpeed = -ballYSpeed
		}

		MyCircle_setX(ball, MyCircle_getX(ball)+ballXSpeed)
		MyCircle_setY(ball, MyCircle_getY(ball)+ballYSpeed)

		MyWindow_clear(w)

		MyRect_display(racket, w)
		MyCircle_display(ball, w)

		for i := 0; i < 6; i++ {
			for j := 0; j < 10; j++ {
				if wall[i][j] {
					MyRect_setPosition(brick, float32(j*100), float32(i*50))
					MyRect_display(brick, w)
				}
			}
		}

		MyWindow_display(w)
	}
}

func init() { runtime.LockOSThread() }
