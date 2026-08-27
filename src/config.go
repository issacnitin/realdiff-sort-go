package sortdemo

import "sort"

func byPriority(rules []DiscountRule) []DiscountRule {
	ordered := append([]DiscountRule(nil), rules...)
	sort.Slice(ordered, func(left, right int) bool {
		return ordered[left].Priority < ordered[right].Priority
	})
	return ordered
}
