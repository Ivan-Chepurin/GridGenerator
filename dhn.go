package main

import "math"

// HDFP height difference from point
type HDFP struct {
	North     float64 `json:"north,omitempty"`
	South     float64 `json:"south,omitempty"`
	East      float64 `json:"east,omitempty"`
	West      float64 `json:"west,omitempty"`
	NorthEast float64 `json:"north_east"`
	NorthWest float64 `json:"north_west"`
	SouthEast float64 `json:"south_east"`
	SouthWest float64 `json:"south_west"`
}

func NewHDFPbyPointSlice(h float64, s [][]float64) *HDFP {
	d := HDFP{
		North:     math.Max(h-s[0][1], s[0][1]-h),
		South:     math.Max(h-s[2][1], s[2][1]-h),
		East:      math.Max(h-s[1][2], s[1][2]-h),
		West:      math.Max(h-s[1][0], s[1][0]-h),
		NorthEast: math.Max(h-s[0][2], s[0][2]-h),
		NorthWest: math.Max(h-s[0][0], s[0][0]-h),
		SouthEast: math.Max(h-s[2][2], s[2][2]-h),
		SouthWest: math.Max(h-s[2][0], s[2][0]-h),
	}
	return &d
}

func NewHDFPbyGridPoint(g *Grid, x, y int) *HDFP {
	d := HDFP{
		North:     math.Max(g.GetHC(x, y)-g.GetHC(x-1, y), g.GetHC(x-1, y)-g.GetHC(x, y)),
		South:     math.Max(g.GetHC(x, y)-g.GetHC(x+1, y), g.GetHC(x+1, y)-g.GetHC(x, y)),
		East:      math.Max(g.GetHC(x, y)-g.GetHC(x, y+1), g.GetHC(x, y+1)-g.GetHC(x, y)),
		West:      math.Max(g.GetHC(x, y)-g.GetHC(x, y-1), g.GetHC(x, y-1)-g.GetHC(x, y)),
		NorthEast: math.Max(g.GetHC(x, y)-g.GetHC(x+1, y-1), g.GetHC(x+1, y-1)-g.GetHC(x, y)),
		NorthWest: math.Max(g.GetHC(x, y)-g.GetHC(x-1, y-1), g.GetHC(x-1, y-1)-g.GetHC(x, y)),
		SouthEast: math.Max(g.GetHC(x, y)-g.GetHC(x+1, y+1), g.GetHC(x+1, y+1)-g.GetHC(x, y)),
		SouthWest: math.Max(g.GetHC(x, y)-g.GetHC(x-1, y+1), g.GetHC(x-1, y+1)-g.GetHC(x, y)),
	}
	return &d
}

func (hd *HDFP) SetNorthMax(n float64) {
	hd.North = math.Max(hd.North, n)
}

func (hd *HDFP) SetNorthMin(n float64) {
	hd.North = math.Min(hd.North, n)
}

func (hd *HDFP) SetSouthMax(n float64) {
	hd.South = math.Max(hd.South, n)
}

func (hd *HDFP) SetSouthMin(n float64) {
	hd.South = math.Min(hd.South, n)
}

func (hd *HDFP) SetEastMax(n float64) {
	hd.East = math.Max(hd.East, n)
}

func (hd *HDFP) SetEastMin(n float64) {
	hd.East = math.Min(hd.East, n)
}

func (hd *HDFP) SetWestMax(n float64) {
	hd.West = math.Max(hd.West, n)
}

func (hd *HDFP) SetWestMin(n float64) {
	hd.West = math.Min(hd.West, n)
}

func (hd *HDFP) SetNorthEastMax(n float64) {
	hd.NorthEast = math.Max(hd.NorthEast, n)
}

func (hd *HDFP) SetNorthEastMin(n float64) {
	hd.NorthEast = math.Min(hd.NorthEast, n)
}

func (hd *HDFP) SetNorthWestMax(n float64) {
	hd.NorthWest = math.Max(hd.NorthWest, n)
}

func (hd *HDFP) SetNorthWestMin(n float64) {
	hd.NorthWest = math.Min(hd.NorthWest, n)
}

func (hd *HDFP) SetSouthEastMax(n float64) {
	hd.SouthEast = math.Max(hd.SouthEast, n)
}

func (hd *HDFP) SetSouthEastMin(n float64) {
	hd.SouthEast = math.Min(hd.SouthEast, n)
}

func (hd *HDFP) SetSouthWestMax(n float64) {
	hd.SouthWest = math.Max(hd.SouthWest, n)
}

func (hd *HDFP) SetSouthWestMin(n float64) {
	hd.SouthWest = math.Min(hd.SouthWest, n)
}

func (hd *HDFP) CompareMax(hd2 *HDFP) {
	hd.SetNorthMax(hd2.North)
	hd.SetSouthMax(hd2.South)
	hd.SetEastMax(hd2.East)
	hd.SetWestMax(hd2.West)
	hd.SetNorthEastMax(hd2.NorthEast)
	hd.SetNorthWestMax(hd2.NorthWest)
	hd.SetSouthEastMax(hd2.SouthEast)
	hd.SetSouthWestMax(hd2.SouthWest)
}

func (hd *HDFP) CompareMin(hd2 *HDFP) {
	hd.SetNorthMin(hd2.North)
	hd.SetSouthMin(hd2.South)
	hd.SetEastMin(hd2.East)
	hd.SetWestMin(hd2.West)
	hd.SetNorthEastMin(hd2.NorthEast)
	hd.SetNorthWestMin(hd2.NorthWest)
	hd.SetSouthEastMin(hd2.SouthEast)
	hd.SetSouthWestMin(hd2.SouthWest)
}
