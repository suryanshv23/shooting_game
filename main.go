package main

import (
	"fmt"
	"go-sdl2-master/sdl"
)

const (
	screenWidth  = 800
	screenHieght = 600
)

func texturefromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading player sprite : %s", err))
	}
	defer img.Free()

	//image goes on RAM but a texture directly goes to graphics card or video card
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Creating player texture : %s", err))
	}

	return tex
}
func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Error : %s", err)
		return
	}
	window, err := sdl.CreateWindow(
		"Gaming in GO",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHieght,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Printf("Creating a new window : %s", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Creating a renderer for window : %s", err)
	}
	defer renderer.Destroy()

	plr := newplayer(renderer)
	if err != nil {
		fmt.Println("Creating Player : ", err)
	}

	var enemies []*enemy

	for i := 0; i < 8; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)*(screenWidth/8) + enemyWidth/2
			y := float64(j)*enemyHeight + enemyHeight/2
			enemy := newenemy(renderer, x, y)
			enemy.active = true
			if err != nil {
				fmt.Println("Creating new enemy", err)
			}
			enemies = append(enemies, &enemy)
		}
	}

	populatebullet(renderer)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		plr.update()

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		plr.draw(renderer)

		checkcollision(bulletpool, enemies)

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}
		for _, bul := range bulletpool {
			bul.draw(renderer)
			bul.update()

		}

		renderer.Present()
	}
}
