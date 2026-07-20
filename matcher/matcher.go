package matcher

import (
	"fmt"
	"math"

	"sahzam/fingerprint"
)

// MatchResult contains the best match information.
type MatchResult struct {
	Song        string
	Similarity float64
	Confidence string
}

// Matcher compares a fingerprint against a database.
type Matcher interface {
	Match(query *fingerprint.Fingerprint, db []fingerprint.FingerprintRecord) (*MatchResult, error)
}

// SimpleMatcher performs a straightforward comparison using shared top-frequency peaks.
type SimpleMatcher struct{}

// Match compares the top frequencies from the query fingerprint to each database fingerprint.
func (m *SimpleMatcher) Match(query *fingerprint.Fingerprint, db []fingerprint.FingerprintRecord) (*MatchResult, error) {
	if query == nil {
		return nil, fmt.Errorf("query fingerprint is nil")
	}
	if len(query.TopFrequencies) == 0 {
		return nil, fmt.Errorf("query fingerprint has no peaks")
	}
	if len(db) == 0 {
		return nil, fmt.Errorf("fingerprint database is empty")
	}

	bestScore := -1.0
	bestSong := ""
	for _, record := range db {
		if record.Fingerprint == nil || len(record.Fingerprint.TopFrequencies) == 0 {
			continue
		}
		score := similarity(query.TopFrequencies, record.Fingerprint.TopFrequencies)
		if score > bestScore {
			bestScore = score
			bestSong = record.Song
		}
	}

	if bestSong == "" {
		return nil, fmt.Errorf("no usable fingerprints found in database")
	}

	confidence := confidenceLabel(bestScore)
	return &MatchResult{Song: bestSong, Similarity: bestScore, Confidence: confidence}, nil
}

func similarity(a, b []int) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	matches := 0
	for _, left := range a {
		for _, right := range b {
			if left == right {
				matches++
				break
			}
		}
	}

	maxLen := float64(len(a))
	if len(b) > len(a) {
		maxLen = float64(len(b))
	}
	return math.Round((float64(matches)/maxLen)*1000) / 10
}

func confidenceLabel(score float64) string {
	switch {
	case score >= 80:
		return "High"
	case score >= 60:
		return "Medium"
	default:
		return "Low"
	}
}
