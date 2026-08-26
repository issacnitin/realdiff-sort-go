package sortdemo

const breakPriorityTiesByCode = false

func configurationBoundary[T any](value T) T {
	return value
}
