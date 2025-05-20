package twopointers_test

import (
	"dsa/twopointers"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	got := twopointers.MaxProfit(prices)

	if expected != got {
		t.Fatalf("TestMaxProfit failed, expected: %v, got: %v", expected, got)
	}
}
