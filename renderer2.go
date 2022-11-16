package main

import "image/color"

func (r *Renderer) GetPixelR(x, y int) *color.RGBA {
	h := r.g.GetHC(x, y)
	if h <= r.maxWaterH {
		return r.WaterPix(h)
	}

	return nil
}
