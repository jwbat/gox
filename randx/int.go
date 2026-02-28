package randx

import "math/rand/v2"


func Int(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return a + rand.IntN(b - a + 1)
}
