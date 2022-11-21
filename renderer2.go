package main

import (
	"image/color"
	"log"
	"math"
)

type Settings struct {
	maxH            float64 // максимальная возможная высота на карте в метрах
	sideLength      float64 // длина стороны карты в метрах
	maxWaterH       float64
	maxShoreH       float64
	maxForestH      float64
	maxMeadowH      float64
	maxMountH       float64
	snow            bool
	pixDist         float64
	diagonalPixDist float64

	g  *Grid      // рассчитанная карта высот
	gi *GridImage // image
}

func (s *Settings) GetRealH(h float64) float64 {
	rh := s.maxH / (s.g.MaxH - s.g.MinH) * (h - s.g.MinH)
	return rh
}

func (s *Settings) SetPixelDistance() {
	s.pixDist = s.sideLength / float64(s.g.Length)
}

func (s *Settings) SetDiagonalPixDist() {
	s.diagonalPixDist = s.pixDist * math.Sqrt(2)
}

func (s *Settings) SurfaceSlope(x, y int) {

}

type RealRenderer struct {
	s Settings
}

func (r *Renderer) GetPixelR(x, y int) *color.RGBA {
	h := r.g.GetHC(x, y)
	if h <= r.maxWaterH {
		return r.WaterPix(h)
	}

	return nil
}

func (r *Renderer) GetPixelH(x, y int, g *Grid) *color.RGBA {
	if r.HorizontalLine(g.GetHC(x, y)) {
		return r.BlackPoint()
	}
	hdfp := g.GetHDFP(x, y)
	bc := r.GetBiomColor(hdfp)
	log.Println(bc)
	return r.BlackPoint()
}

func (r *Renderer) GetBiomColor(hdfp *HDFP) *color.RGBA {

	return nil
}
