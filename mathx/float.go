package mathx

import "math"


func Round(x float64, places int) float64 {
	scale := math.Pow(10, float64(places))
	return math.Round(x * scale) / scale
}
