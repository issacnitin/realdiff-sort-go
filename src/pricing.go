package sortdemo

type DiscountRule struct {
	Code         string
	Priority     int
	MinimumTotal int
	PercentOff   int
}

func byPriority(rules []DiscountRule) []DiscountRule {
	ordered := append([]DiscountRule(nil), rules...)
	sortRules(ordered, func(left, right int) bool {
		return ordered[left].Priority < ordered[right].Priority
	})
	return ordered
}

func selectDiscount(listPrice int) DiscountRule {
	rules := []DiscountRule{
		{Code: "INELIGIBLE_00", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_01", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_02", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_03", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_04", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "Z_CLEARANCE", Priority: 3, MinimumTotal: 50, PercentOff: 40},
		{Code: "INELIGIBLE_06", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_07", Priority: 1, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_08", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_09", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_10", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_11", Priority: 0, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_12", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_13", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_14", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_15", Priority: 0, MinimumTotal: 1000, PercentOff: 5},
		{Code: "A_SEASONAL", Priority: 3, MinimumTotal: 50, PercentOff: 15},
		{Code: "INELIGIBLE_17", Priority: 0, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_18", Priority: 0, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_19", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_20", Priority: 0, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_21", Priority: 1, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_22", Priority: 3, MinimumTotal: 1000, PercentOff: 5},
		{Code: "INELIGIBLE_23", Priority: 2, MinimumTotal: 1000, PercentOff: 5},
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
