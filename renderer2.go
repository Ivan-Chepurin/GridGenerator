package main

import (
	"image/color"
	"log"
	"math"
)

type Settings struct {
	MaxH               float64 `json:"max_h"`             // максимальная возможная высота на карте в метрах над уровнем моря
	MaxWaterDepth      float64 `json:"max_water_depth"`   // максимальная глубина моря
	SideLength         float64 `json:"side_length"`       // длина стороны карты в метрах
	PixDist            float64 `json:"pix_dist"`          // дистанция между соседними пикселями
	DiagonalPixDist    float64 `json:"diagonal_pix_dist"` // дистанция между пикселями по диагонали
	MaxWaterH          float64 `json:"max_water_h"`       // высота воды в процентах от всей высоты
	MaxShoreH          float64 `json:"max_shore_h"`
	MaxForestH         float64 `json:"max_forest_h"`
	MaxMeadowH         float64 `json:"max_meadow_h"`
	MaxMountH          float64 `json:"max_mount_h"`
	Snow               bool    `json:"snow"`
	HorizontalInterval int64   `json:"horizontal_interval"` // интервал между линиями высот (x*1/350)

	g *Grid // рассчитанная карта высот
}

type RealRenderer struct {
	s Settings

	alpha uint8 // прозрачность

	gi *GridImage // image
}

func NewRealRenderer(s Settings, gi *GridImage) *RealRenderer {
	return &RealRenderer{
		s:     s,
		alpha: 0xff,
		gi:    gi,
	}
}

func (s *Settings) GetRealH(h float64) float64 {
	rh := (s.MaxH + s.MaxWaterDepth) / (s.g.MaxH - s.g.MinH) * (h - s.g.MinH)
	return rh
}

func (s *Settings) SetPixelDistance() {
	s.PixDist = s.SideLength / float64(s.g.Length)
}

func (s *Settings) SetDiagonalPixDist() {
	s.DiagonalPixDist = s.PixDist * math.Sqrt(2)
}

func (s *Settings) SurfaceSlope(x, y int) {

}

func (r *RealRenderer) GetPixel(x, y int) *color.RGBA {
	h := r.s.GetRealH(r.s.g.GetHC(x, y))
	if h <= r.s.MaxWaterH {
		return r.s.WaterPix(h)
	}

	return nil
}

func (r *RealRenderer) Render() {
	var cl []chan int
	for sy := 0; sy <= r.s.g.Width-1; sy += r.s.g.Width / 4 {
		for sx := 0; sx <= r.s.g.Length-1; sx += r.s.g.Length / 4 {
			c := make(chan int)
			cl = append(cl, c)
			go r.RenderRegion(sx, sx+r.s.g.Length/4, sy, sy+r.s.g.Width/4, c)
		}
	}
	for c := range cl {
		<-cl[c]
		close(cl[c])
	}
}

func (r *RealRenderer) RenderRegion(sx, fx, sy, fy int, c chan int) {
	for y := sy; y <= fy; y++ {
		for x := sx; x <= fx; x++ {
			r.gi.SetPoint(x, y, r.GetPixel(x, y))
		}
	}
	c <- 1
}

func (r *RealRenderer) GetPixelH(x, y int) *color.RGBA {
	if r.HorizontalLine(r.s.g.GetHC(x, y)) {
		return r.BlackPoint()
	}
	hd := NewHeightDifference(&r.s, x, y)
	bc := r.GetBiomColor(hd)
	log.Println(bc)
	return r.BlackPoint()
}

func (r *RealRenderer) GetBiomColor(hd *HeightDifference) *color.RGBA {

	return nil
}

func (s *Settings) WaterPix(h float64) *color.RGBA {
	// to do
	return nil
}

func (r *RealRenderer) HorizontalLine(h float64) bool {
	if r.s.HorizontalInterval < 350 {
		return int64(h/((r.s.g.MaxH-r.s.g.MinH)/350))%r.s.HorizontalInterval == 0
	}
	return false
}

func (r *RealRenderer) BlackPoint() *color.RGBA {
	return &color.RGBA{A: r.alpha}
}
