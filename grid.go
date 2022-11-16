package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func average(a, b, c, d float64) float64 {
	return (a + b + c + d) / 4
}

type Grid struct {
	Roughness                 float64 `json:"Roughness,omitempty"`
	Size                      int     `json:"Size,omitempty"`
	Length                    int     `json:"length,omitempty"`
	Width                     int     `json:"Width,omitempty"`
	MaxH                      float64 `json:"MaxH,omitempty"`
	MinH                      float64 `json:"MinH,omitempty"`
	HeightReductionPercentage int     `json:"height_reduction_percentage"`
	PowOn                     bool    `json:"pow_on"`

	grid []float64
	mu   sync.Mutex
}

func NewGrid(roughness float64, l, w, hrp int, powOn bool) *Grid {
	g := &Grid{
		Roughness:                 roughness,
		Size:                      l * w / 10000,
		Length:                    l + 1,
		Width:                     w + 1,
		MinH:                      float64(w * l),
		HeightReductionPercentage: hrp,
		PowOn:                     powOn,
	}
	var cells []float64
	for i := 0; i <= g.Width*g.Length; i++ {
		cells = append(cells, 0)
	}
	g.grid = cells
	g.MakeLandscape()
	g.PowLandscape()
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	return g
}

func (g *Grid) RandHeight(size int) float64 {
	r := g.Roughness * float64(rand.Intn(size)+1) * rand.Float64()
	if rand.Intn(100)+1 > g.HeightReductionPercentage {
		return r
	}
	return -r

}

func (g *Grid) SetMaxH(h float64) {
	if h > g.MaxH {
		g.MaxH = h
	}
}

func (g *Grid) SetMinH(h float64) {
	if h < g.MinH {
		g.MinH = h
	}
}

func (g *Grid) GetHC(x, y int) float64 {
	if x < 0 || x > g.Length-1 || x == 0 || x == g.Length-1 ||
		y < 0 || y > g.Width-1 || y == 0 || y == g.Width-1 {
		return 10 + float64(rand.Intn(100))
	}
	return g.grid[x+g.Length*y]
}

func (g *Grid) SetH(x, y int, h float64) {
	if h < 0 {
		h *= -1
	}
	g.grid[x+g.Length*y] = h
}

func (g *Grid) MakeLandscape() {
	rand.Seed(time.Now().Unix())
	g.SetStartPoints()
	g.divide(g.Length / 2)
}

func (g *Grid) SetStartPoints() {
	g.SetH(0, 0, 3*rand.Float64())
	g.SetH(g.Length, 0, 3*rand.Float64())
	g.SetH(0, g.Width, 3*rand.Float64())
	g.SetH(g.Length-1, g.Width-1, 3*rand.Float64())
	g.SetH(g.Length/2, g.Width/2, float64(rand.Intn(g.Size)))
}

func (g *Grid) divide(size int) {
	half := size / 2
	x := half
	y := x
	if float64(size)/2 < 1 {
		return
	}

	for y = half; y <= g.Width-1; y += size {
		for x = half; x < g.Length-1; x += size {
			g.square(x, y, half, g.RandHeight(size))
		}
	}
	for y = 0; y <= g.Width-1; y += half {
		for x = (y + half) % size; x <= g.Length-1; x += size {
			g.diamond(x, y, half, g.RandHeight(size))
		}
	}
	g.divide(size / 2)
}

func (g *Grid) square(x, y, size int, offset float64) {
	ave := average(
		g.GetHC(x-size, y-size),
		g.GetHC(x+size, y+size),
		g.GetHC(x-size, y+size),
		g.GetHC(x+size, y-size),
	)
	g.SetH(x, y, ave+offset)
}

func (g *Grid) diamond(x, y, size int, offset float64) {
	ave := average(
		g.GetHC(x, y+size),
		g.GetHC(x, y-size),
		g.GetHC(x+size, y),
		g.GetHC(x-size, y),
	)
	g.SetH(x, y, ave+offset)
}

func (g *Grid) PowLandscape() {
	for i := 0; i < len(g.grid); i++ {
		if g.PowOn {
			g.grid[i] = math.Pow(g.grid[i], 2)
		}
		g.SetMaxH(g.grid[i])
		g.SetMinH(g.grid[i])
	}
}
