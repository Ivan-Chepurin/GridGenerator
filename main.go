package main

import (
	"math"
)

func main() {
	coef := 0.22
	w := int(math.Pow(2, 12))
	g := NewGrid(1.9, w, w, 30)
	r := NewRenderer(
		g, coef, 0.02, 0.4, 0.4, 0.6, true, 50,
	)
	r.Render()
	r.saveImage("grid")
}
