package mymath

import (
	"fmt"
	"regexp"
	"strconv"
)

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

func FloatPrec(f float64, min, max int) (string, error) {
	if min >= max {
		return "", fmt.Errorf("min must be < max in call to FloatPrec")
	}
	str := strconv.FormatFloat(f, 'f', max+1, 64)
	pattern := fmt.Sprintf("^\\d+\\.\\d{%d}(\\d{0,%d}[1-9])?", min, max-min-1)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	match := re.FindString(str)
	return match, nil
}
