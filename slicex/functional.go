package slicex


import "cmp"


func Map[T1, T2 any](s []T1, f func (T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}


func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}


func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}


type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}


func Sum[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}


func Mul[T Number](s []T) T {
	var product T = 1
	for _, v := range s {
		product *= v
	}
	return product
}


func Max[T cmp.Ordered](s []T) (T, bool) {
    if len(s) == 0 {
        var zero T
        return zero, false
    }
    m := s[0]
    for _, v := range s[1:] {
        if v > m {
            m = v
        }
    }
    return m, true
}


func Min[T cmp.Ordered](s []T) (T, bool) {
    if len(s) == 0 {
        var zero T
        return zero, false
    }
    m := s[0]
    for _, v := range s[1:] {
        if v < m {
            m = v
        }
    }
    return m, true
}
