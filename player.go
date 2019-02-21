package main

import (
	"go-sdl2-master/sdl"
	"math"
	"time"
)

const (
	playerSpeed  = 0.45
	playerHeight = 150
	playerWidth  = 157
	cooldown     = time.Millisecond * 250
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastshot time.Time
}

func newplayer(renderer *sdl.Renderer) (p player) {

	p.tex = texturefromBMP(renderer, "Images/rocket.bmp")

	p.x = (screenWidth / 2) - (playerWidth / 2)
	p.y = screenHieght - playerHeight

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 152, H: 150},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 152, H: 150})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x+playerWidth > 0 {
			p.x -= playerSpeed
		} else {
			p.x = screenWidth
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x < screenWidth {
			p.x += playerSpeed
		} else {
			p.x = -playerWidth
		}
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastshot) >= cooldown {
			p.shoot()
			p.lastshot = time.Now()
		}
	}
}

func (p *player) shoot() {
	if bul, ok := popbullet(); ok {
		bul.active = true
		bul.x = p.x
		bul.y = p.y
		bul.angle = 270 * (math.Pi / 180)
		p.lastshot = time.Now()
	}
}
