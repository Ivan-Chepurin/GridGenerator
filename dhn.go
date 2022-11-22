package main

import "math"

// HeightDifference height difference from point
//type HeightDifference struct {
//	NAngle   float64 `json:"north,omitempty"`
//	SAngle   float64 `json:"south,omitempty"`
//	EAngle   float64 `json:"east,omitempty"`
//	WAngle   float64 `json:"west,omitempty"`
//	NEAngle  float64 `json:"north_east"`
//	NWAngle  float64 `json:"north_west"`
//	SEAngle  float64 `json:"south_east"`
//	SWAngle  float64 `json:"south_west"`
//	MaxAngle float64 `json:"max_angle"`
//	LiftDir  int     `json:"lift_dir"`
//}

type HeightDifference struct {
	MaxAngle   float64 `json:"max_angle"`
	LiftDir    int     `json:"lift_dir"`
	MinAngle   float64 `json:"min_angle"`
	DescentDir int     `json:"descent_dir"`
	TopPoint   bool    `json:"top_point"`
	LowPoint   bool    `json:"low_point"`
}

func GetAngle(h1, h2, dist float64) float64 {
	max := math.Max(h1, h2)
	min := math.Min(h1, h2)
	bc := max - min
	atan := math.Atan(bc / dist)
	if h1 < h2 {
		return atan
	}
	return -1 * atan
}

func NewHeightDifference(set *Settings, x, y int) *HeightDifference {
	angles := make([]float64, 8)
	pointH := set.g.GetHC(x, y)
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x-1, y), set.PixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x+1, y), set.PixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x, y+1), set.PixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x, y-1), set.PixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x+1, y-1), set.DiagonalPixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x-1, y-1), set.DiagonalPixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x+1, y+1), set.DiagonalPixDist))
	angles = append(angles, GetAngle(pointH, set.g.GetHC(x-1, y+1), set.DiagonalPixDist))

	maxAngle := 0.0
	indexMaxAngle := 0
	minAngle := 0.0
	indexMinAngle := 0
	lowPoint := false
	topPoint := false
	for index, angle := range angles {
		if maxAngle < angle {
			maxAngle = angle
			indexMaxAngle = index
		}
		if minAngle > angle {
			minAngle = angle
			indexMinAngle = index
		}
	}
	if maxAngle == 0.0 {
		indexMaxAngle = 9
		topPoint = true
	}
	if minAngle == 0.0 {
		indexMinAngle = 8
		lowPoint = true
	}
	hd := HeightDifference{
		MaxAngle:   maxAngle,
		LiftDir:    indexMaxAngle,
		MinAngle:   minAngle,
		DescentDir: indexMinAngle,
		TopPoint:   topPoint,
		LowPoint:   lowPoint,
	}
	return &hd
}
