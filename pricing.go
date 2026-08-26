package sortdemo

import "sort"

type DiscountRule struct {
	Code         string
	Priority     int
	MinimumTotal int
	PercentOff   int
}

func byPriority(rules []DiscountRule) []DiscountRule {
	ordered := append([]DiscountRule(nil), rules...)
	sort.SliceStable(ordered, func(left, right int) bool {
		if breakPriorityTiesByCode && ordered[left].Priority == ordered[right].Priority {
			return ordered[left].Code < ordered[right].Code
		}
		return ordered[left].Priority < ordered[right].Priority
	})
	return ordered
}

func selectDiscount(listPrice int) DiscountRule {
	rules := []DiscountRule{
		{Code: "Z_CLEARANCE", Priority: 10, MinimumTotal: 50, PercentOff: 40},
		{Code: "A_SEASONAL", Priority: 10, MinimumTotal: 50, PercentOff: 15},
		{Code: "INELIGIBLE", Priority: 10, MinimumTotal: 1000, PercentOff: 5},
	}
	for _, rule := range byPriority(rules) {
		if listPrice >= rule.MinimumTotal {
			return rule
		}
	}
	panic("no eligible discount")
}

func checkoutTotal(listPrice int) (int, string) {
	selected := selectDiscount(listPrice)
	return listPrice * (100 - selected.PercentOff) / 100, selected.Code
}

func exerciseCoverage() int {
	return exerciseComparisonBreadth()
}
