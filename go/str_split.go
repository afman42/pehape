package pehape

import (
	"errors"
)

var (
	errMust2Params         = errors.New("expects at most 2 parameters")
	errMustGreaterThanZero = errors.New("The length of each segment must be greater than zero")
)

func StrSplit(s string, length ...int) ([]string, error) {
	if len(length) > 1 {
		return nil, errMust2Params
	}

	leng := 1
	if len(length) == 1 {
		leng = length[0]
	}

	if leng < 1 {
		return nil, errMustGreaterThanZero
	}

	lenS := len(s)
	if lenS == 0 {
		return []string{""}, nil
	}

	if leng > lenS {
		return []string{s}, nil
	}

	capacity, mod := getCapacity(lenS, leng)

	results := make([]string, 0, capacity)

	to := leng
	for i := 0; i < lenS; i += leng {
		if to <= lenS {
			results = append(results, s[i:to])
			to += leng
		}
	}

	if mod > 0 {
		results = append(results, s[lenS-mod:])
	}
	return results, nil
}

func getCapacity(lenS, n int) (int, int) {
	mod := lenS % n
	if mod == 0 {
		return lenS / n, 0
	}

	return lenS/n + 1, mod
}
