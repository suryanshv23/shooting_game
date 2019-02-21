package main

import "math"

type circle struct {
	x, y   float64
	radius float64
}

func iscollision(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow((c1.x-c2.x), 2) + math.Pow((c1.y-c2.y), 2))

	return dist <= c1.radius+c2.radius
}

func checkcollision(bulletpool []*bullet, enemies []*enemy) {
	for _, i := range bulletpool {
		if i.active {
			for _, j := range enemies {
				if j.active {
					c1 := circle{x: ((i.x + bulletwidth) / 2), y: i.y - bulletheight/2, radius: bulletheight / 2}
					c2 := circle{x: ((j.x + enemyWidth) / 2) - 10, y: j.y - enemyHeight/2, radius: enemyHeight / 2}
					res := iscollision(c1, c2)
					if res == true {
						i.active = false
						j.active = false
					}
				}
			}
		}
	}
}
