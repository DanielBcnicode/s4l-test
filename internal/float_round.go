package internal

import "math"

// FloatRoundPrecision return a float32 value with the rounded p decimals
func FloatRoundPrecision(n float32, p int) float32 {
	return float32(math.Round(float64(n)*(math.Pow10(p))) / math.Pow10(p))
}
