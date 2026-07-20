package fingerprint

import "fmt"

// Fingerprint is a compact representation of an audio clip.
type Fingerprint struct {
	TopFrequencies []int
	WindowSize     int
}

// FingerprintRecord pairs a song with its fingerprint.
type FingerprintRecord struct {
	Song        string
	Fingerprint *Fingerprint
}

// Generator creates fingerprints from audio samples.
type Generator interface {
	Generate(samples []float64, sampleRate int) (*Fingerprint, error)
}

// ErrEmptySignal indicates that there are no samples to process.
var ErrEmptySignal = fmt.Errorf("empty signal")

// ErrInvalidWindow indicates an invalid window size.
var ErrInvalidWindow = fmt.Errorf("invalid window size")
