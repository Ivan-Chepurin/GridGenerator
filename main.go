package main

import (
	"math"
)

func main() {
	w := int(math.Pow(2, 12))
	g := NewGrid(1.9, w, w, 30, false)
	r := NewRenderer(
		g, 0.32, 0.02, 0.2, 0.15, 0.2, 528,
	)
	r.Render()
	r.saveImage("grid")
}
