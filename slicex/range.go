package slicex


func Range(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i  
	}
	return r
}


func RangeStep(start, stop, step int) []int {
	// [start, stop] inclusive
	if step == 0 || (step > 0 && start >= stop) || (step < 0 && start <= stop) {
		return []int{}
	}
	d := stop - start
	s := step
	n := d / s + 1
	r := make([]int, 0, n)
	for i := start; (step > 0 && i <= stop) || (step < 0 && i >= stop); i += step {
		r = append(r, i)
	}
	return r
}
