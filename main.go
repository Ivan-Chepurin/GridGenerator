package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	w := int(math.Pow(2, 12))
	g := NewGrid(3.9, w, w, 39, false)
	r := NewRenderer(
		g, 0.34, 0.01, 0.2, 0.15, 0.2, 528,
	)
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	r.Render()
	r.saveImage("grid")
}
