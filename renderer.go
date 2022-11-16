package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

type GridImage struct {
	i  *image.RGBA
	mu sync.Mutex
}

func (gi *GridImage) SetPoint(x, y int, p *color.RGBA) {
	gi.i.Set(x, y, p)
}

type Renderer struct {
	maxWaterH  float64 // вода
	maxShoreH  float64 // берег
	maxForestH float64 // лес
	maxMeadowH float64 // луга
	maxMountH  float64 // горы
	snow       bool

	horizontalInterval int64 // интервал между линиями высот (x*1/350)
	alpha              uint8 // прозрачность

	g   *Grid      // рассчитанная карта высот
	gic *GridImage // image
	gir *GridImage // image
}

func NewRenderer(
	grid *Grid, waterHCoef, shoreH, forestH, meadowH, mountH float64, horizontalInterval int64,
) *Renderer {

	maxWaterH := grid.MinH + (grid.MaxH-grid.MinH)*waterHCoef
	maxShoreH := maxWaterH + (grid.MaxH-grid.MinH)*shoreH
	maxForestH := maxShoreH + (grid.MaxH-grid.MinH)*forestH
	maxMeadowH := maxForestH + (grid.MaxH-grid.MinH)*meadowH
	maxMountH := maxMeadowH + (grid.MaxH-grid.MinH)*mountH

	fmt.Printf(
		"maxWaterH := %v\n"+
			"maxShoreH := %v\n"+
			"maxForestH := %v\n"+
			"maxMeadowH := %v\n"+
			"maxMountH := %v\n",
		maxWaterH,
		maxShoreH,
		maxForestH,
		maxMeadowH,
		maxMountH,
	)

	r := Renderer{
		maxWaterH:          maxWaterH,
		maxShoreH:          maxShoreH,
		maxForestH:         maxForestH,
		maxMeadowH:         maxMeadowH,
		maxMountH:          maxMountH,
		horizontalInterval: horizontalInterval,
		alpha:              0xff,
		g:                  grid,
		gic: &GridImage{
			i: image.NewRGBA(image.Rectangle{
				Min: image.Point{},
				Max: image.Point{X: grid.Length, Y: grid.Width},
			}),
		},
	}
	return &r
}

func (r *Renderer) Render() {
	var cl []chan int
	for sy := 0; sy <= r.g.Width-1; sy += r.g.Width / 4 {
		for sx := 0; sx <= r.g.Length-1; sx += r.g.Length / 4 {
			c := make(chan int)
			cl = append(cl, c)
			go r.RenderRegion(sx, sx+r.g.Length/4, sy, sy+r.g.Width/4, c)
		}
	}
	for c := range cl {
		<-cl[c]
		close(cl[c])
	}
}

func (r *Renderer) RenderRegion(sx, fx, sy, fy int, c chan int) {
	for y := sy; y <= fy; y++ {
		for x := sx; x <= fx; x++ {
			r.gic.SetPoint(x, y, r.GetPixelC(r.g.GetHC(x, y)))
			//r.gir.SetPoint(x, y, r.GetPixelR(x, y))
		}
	}
	c <- 1
}

func (r *Renderer) HorizontalLine(h float64) bool {
	if r.horizontalInterval < 350 {
		return int64(h/((r.g.MaxH-r.g.MinH)/350))%r.horizontalInterval == 0
	}
	return false
}

func (r *Renderer) GetPixelC(h float64) *color.RGBA {
	if r.HorizontalLine(h) {
		return &color.RGBA{A: r.alpha}
	}
	if h <= r.maxWaterH {
		return r.WaterPix(h)
	}
	if h > r.maxWaterH && h <= r.maxShoreH {
		return r.ShorePix(h)
	}
	if h > r.maxShoreH && h <= r.maxForestH {
		return r.ForestPix(h)
	}
	if h > r.maxForestH && h <= r.maxMeadowH {
		return r.MeadowPix(h)
	}
	if h > r.maxMeadowH && h <= r.maxMountH {
		return r.MountainPix(h)
	}
	if h > r.maxMountH {
		return r.SnowPix(h)
	}
	fmt.Println("ERROR")
	return &color.RGBA{A: r.alpha}
}

func (r *Renderer) ShorePix(h float64) *color.RGBA {
	return &color.RGBA{
		R: uint8(218 + (236-218)/(r.maxShoreH-r.maxWaterH)*(h-r.maxWaterH)),
		G: uint8(165 + (210-165)/(r.maxShoreH-r.maxWaterH)*(h-r.maxWaterH)),
		B: uint8(122 + (190-122)/(r.maxShoreH-r.maxWaterH)*(h-r.maxWaterH)),
		A: r.alpha,
	}
}

func (r *Renderer) WaterPix(h float64) *color.RGBA {
	return &color.RGBA{
		R: uint8(10 + (40-10)/(r.maxWaterH-r.g.MinH)*(h-r.g.MinH+1)),
		G: uint8(20 + (60-20)/(r.maxWaterH-r.g.MinH)*(h-r.g.MinH+1)),
		B: uint8(70 + (254-70)/(r.maxWaterH-r.g.MinH)*(h-r.g.MinH+1)),
		A: r.alpha}
}

func (r *Renderer) ForestPix(h float64) *color.RGBA {
	return &color.RGBA{
		R: uint8(27 + (55-33)/(r.maxForestH-r.maxShoreH)*(h-r.maxShoreH)),
		G: uint8(71 + (126-87)/(r.maxForestH-r.maxShoreH)*(h-r.maxShoreH)),
		B: uint8(22 + (44-26)/(r.maxForestH-r.maxShoreH)*(h-r.maxShoreH)),
		A: r.alpha,
	}
}

func (r *Renderer) MeadowPix(h float64) *color.RGBA {
	return &color.RGBA{
		R: uint8(176 + (222-176)/(r.maxMeadowH-r.maxForestH)*(h-r.maxForestH)),
		G: uint8(192 + (234-192)/(r.maxMeadowH-r.maxForestH)*(h-r.maxForestH)),
		B: uint8(26 + (109-26)/(r.maxMeadowH-r.maxForestH)*(h-r.maxForestH)),
		A: r.alpha,
	}
}

func (r *Renderer) MountainPix(h float64) *color.RGBA {
	c := uint8(87 + (210-87)/(r.maxMountH-r.maxMeadowH)*(h-r.maxMeadowH))
	return &color.RGBA{R: c, G: c, B: c, A: r.alpha}
}

func (r *Renderer) SnowPix(h float64) *color.RGBA {
	return &color.RGBA{
		R: uint8(211 + (255-241)/(r.g.MaxH-r.maxMountH)*(h-r.maxMountH)),
		G: uint8(212 + (255-242)/(r.g.MaxH-r.maxMountH)*(h-r.maxMountH)),
		B: uint8(243 + (255-243)/(r.g.MaxH-r.maxMountH)*(h-r.maxMountH)),
		A: r.alpha,
	}
}

func (r *Renderer) saveImage(name string) {
	f, _ := os.Create(name + ".png")
	err := png.Encode(f, r.gic.i)
	if err != nil {
		return
	}
}
