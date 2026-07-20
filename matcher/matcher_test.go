package matcher

import (
	"testing"

	"sahzam/fingerprint"
)

func TestSimpleMatcherMatch(t *testing.T) {
	matcher := &SimpleMatcher{}
	query := &fingerprint.Fingerprint{TopFrequencies: []int{440, 880, 1320}}
	db := []fingerprint.FingerprintRecord{{Song: "Song A", Fingerprint: &fingerprint.Fingerprint{TopFrequencies: []int{440, 880, 1320}}}}

	result, err := matcher.Match(query, db)
	if err != nil {
		t.Fatalf("expected match to succeed, got %v", err)
	}
	if result.Song != "Song A" {
		t.Fatalf("expected Song A, got %s", result.Song)
	}
	if result.Similarity <= 0 {
		t.Fatalf("expected positive similarity, got %.2f", result.Similarity)
	}
}
