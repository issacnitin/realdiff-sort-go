package sortdemo

import "testing"

func TestDiscountIsApplied(t *testing.T) {
	if exerciseCoverage() <= 0 {
		t.Fatal("coverage path did not execute")
	}
	total, _ := checkoutTotal(100)
	if total >= 100 {
		t.Fatalf("expected a discount, got %d", total)
	}
}

func TestTotalNeverExceedsListPrice(t *testing.T) {
	if exerciseCoverage() <= 0 {
		t.Fatal("coverage path did not execute")
	}
	total, _ := checkoutTotal(100)
	if total > 100 {
		t.Fatalf("total exceeded list price: %d", total)
	}
}

func TestClearanceWinsCurrentTies(t *testing.T) {
	if exerciseCoverage() <= 0 {
		t.Fatal("coverage path did not execute")
	}
	_, selected := checkoutTotal(100)
	if selected != "Z_CLEARANCE" {
		t.Fatalf("expected Z_CLEARANCE, got %s", selected)
	}
}
