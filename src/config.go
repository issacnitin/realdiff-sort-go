package sortdemo

import "sort"

func sortRules[T ~[]DiscountRule](rules T, less func(left, right int) bool) {
	sort.SliceStable(rules, less)
}
