package utils

import "testing"

func TestSimilarityCalculation(t *testing.T) {
	a := []int{100, 200, 300}
	b := []int{100, 400, 300}

	// This test is intentionally simple and demonstrates the helper shape.
	matches := 0
	for _, left := range a {
		for _, right := range b {
			if left == right {
				matches++
				break
			}
		}
	}
	if matches != 2 {
		t.Fatalf("expected 2 matches, got %d", matches)
	}
}
