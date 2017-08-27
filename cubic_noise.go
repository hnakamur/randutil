package randutil

import "math"

const (
	rndA = 134775813
	rndB = 1103515245
)

type CubicNoise struct {
	seed    int32
	octave  int32
	periodX int32
	periodY int32
}

func NewCubicNoise(seed, octave, periodX, periodY int32) *CubicNoise {
	if periodX == 0 {
		periodX = math.MaxInt32
	}
	if periodY == 0 {
		periodY = math.MaxInt32
	}
	return &CubicNoise{
		seed:    seed,
		octave:  octave,
		periodX: periodX,
		periodY: periodY,
	}
}

func (n *CubicNoise) Sample1D(x float64) float64 {
	xi := int32(x) / n.octave
	lerp := float64(x)/float64(n.octave) - float64(xi)

	return n.interpolate(
		n.randomize(n.seed, n.tile(xi-1, n.periodX), 0),
		n.randomize(n.seed, n.tile(xi, n.periodX), 0),
		n.randomize(n.seed, n.tile(xi+1, n.periodX), 0),
		n.randomize(n.seed, n.tile(xi+2, n.periodX), 0),
		lerp)*0.5 + 0.25
}

func (n *CubicNoise) randomize(seed, x, y int32) float64 {
	return float64(((x^y)*rndA^(seed+x))*(((rndB*x)<<16)^(rndB*y)-rndA)) / float64(math.MaxInt32)
}
func (n *CubicNoise) tile(coordinate, period int32) int32 {
	return coordinate % period
}

func (n *CubicNoise) interpolate(a, b, c, d, x float64) float64 {
	p := (d - c) - (a - b)
	return x*x*x*p + x*x*((a-b)-p) + x*(c-a) + b
}
