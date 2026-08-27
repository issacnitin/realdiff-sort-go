package sortdemo

import "sort"

var sortRules = sort.SliceStable

func configurationBoundary[T any](value T) T {
	return value
}
