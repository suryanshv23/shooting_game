package main

import (
	"go-sdl2-master/sdl"
	"math"
)

const (
	bulletheight = 16
	bulletwidth  = 20
	bulletspeed  = 0.6
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
	angle  float64
}

func newbullet(renderer *sdl.Renderer) (bul bullet) {

	bul.tex = texturefromBMP(renderer, "images/bullet.bmp")
	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}
	x := bul.x - bulletwidth/2 + playerWidth/2
	y := bul.y - bulletheight/2

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletwidth, H: bulletheight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletwidth, H: bulletheight})
}

var bulletpool []*bullet

func populatebullet(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newbullet(renderer)
		bulletpool = append(bulletpool, &bul)
	}
}

func popbullet() (*bullet, bool) {
	for _, bul := range bulletpool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}

func (bul *bullet) update() {
	bul.x += bulletspeed * (math.Cos(bul.angle))
	bul.y += bulletspeed * (math.Sin(bul.angle))

	if bul.x > screenWidth || bul.x < -playerWidth || bul.y > screenHieght || bul.y < 0 {
		bul.active = false
	}
}
