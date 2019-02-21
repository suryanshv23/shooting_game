package main

import (
	"go-sdl2-master/sdl"
)

const (
	enemyWidth  = 36
	enemyHeight = 42
)

type enemy struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
}

func newenemy(renderer *sdl.Renderer, x, y float64) (be enemy) {

	be.tex = texturefromBMP(renderer, "Images/enemy.bmp")

	be.x = x
	be.y = y
	return be
}

func (be *enemy) draw(renderer *sdl.Renderer) {
	x := 30 + be.x - enemyWidth/2
	y := be.y - enemyHeight/2
	if be.active == false {
		return
	}

	renderer.CopyEx(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemyWidth, H: enemyHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemyWidth, H: enemyHeight},
		0,
		&sdl.Point{X: enemyWidth / 2, Y: enemyHeight / 2},
		sdl.FLIP_NONE)
}
