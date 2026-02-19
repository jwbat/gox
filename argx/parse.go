package argx

import (
	"fmt"
	"strconv"
)

func ParseInts(args []string) ([]int, error) {
	ns := make([]int, len(args))
	for i, s := range args {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("arg %d (%q): %w", i + 1, s, err)
		}
		ns[i] = n
	}
	return ns, nil
}


func ParseFloat64s(args []string) ([]float64, error) {
	ns := make([]float64, len(args))
	for i, s := range args {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, fmt.Errorf("arg %d (%q): %w", i + 1, s, err)
		}
		ns[i] = f
	}
	return ns, nil
}
