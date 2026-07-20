package fingerprint

import "testing"

func TestFingerprintStructure(t *testing.T) {
	fp := &Fingerprint{TopFrequencies: []int{100, 200, 300}, WindowSize: 1024}
	if len(fp.TopFrequencies) != 3 {
		t.Fatalf("expected 3 top frequencies, got %d", len(fp.TopFrequencies))
	}
	if fp.WindowSize != 1024 {
		t.Fatalf("expected window size 1024, got %d", fp.WindowSize)
	}
}
