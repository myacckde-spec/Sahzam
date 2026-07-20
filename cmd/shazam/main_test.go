package main

import (
	"testing"
)

func TestFingerprintGeneratorGenerate(t *testing.T) {
	generator := &FingerprintGenerator{}

	samples := make([]float64, 2048)
	for i := range samples {
		samples[i] = 8000 * float64(i%2) // a simple alternating signal
	}

	fingerprint, err := generator.Generate(samples, 16000)
	if err != nil {
		t.Fatalf("expected fingerprint generation to succeed, got %v", err)
	}
	if fingerprint == nil {
		t.Fatal("expected non-nil fingerprint")
	}
	if len(fingerprint.TopFrequencies) == 0 {
		t.Fatal("expected at least one frequency peak")
	}
}
