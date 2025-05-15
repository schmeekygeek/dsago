package arrays_test

import (
	"dsa/arrays"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	got := arrays.MaxProfit(prices)

	if expected != got {
		t.Fatalf("TestMaxProfit failed, expected: %v, got: %v", expected, got)
	}
}
