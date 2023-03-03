package internal

import "math"

func FloatRoundPrecision(n float32, p int) float32 {
	return float32(math.Round(float64(n)*(math.Pow10(p))) / math.Pow10(p))
}
