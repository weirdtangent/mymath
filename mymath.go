package mymath

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func MinMax(x, min, max int64) int64 {
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
