package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	w := int(math.Pow(2, 5))
	g := NewGrid(4.8, w, w, 60, false)
	r := NewRenderer(
		g, 0.25, 0.01, 0.2, 0.15, 0.2, 528,
	)
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	r.Render()
	r.saveImage("grid")
}
