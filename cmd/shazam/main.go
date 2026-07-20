package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"sahzam/audio"
	"sahzam/database"
	"sahzam/fft"
	"sahzam/fingerprint"
	"sahzam/matcher"
	"sahzam/reader"
	"sahzam/utils"
)

// App wires the MVP workflow together.
type App struct {
	Reader      audio.Reader
	Generator   fingerprint.Generator
	Matcher     matcher.Matcher
	Database    *database.DemoDatabase
	Logger      utils.Logger
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/shazam <path-to-wav>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		fmt.Printf("resolve path: %v\n", err)
		os.Exit(1)
	}

	app := &App{
		Reader:    &reader.WAVReader{},
		Generator: &FingerprintGenerator{},
		Matcher:   &matcher.SimpleMatcher{},
		Database:  database.NewDemoDatabase(),
		Logger:    &utils.StdLogger{},
	}

	result, err := app.Run(absPath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nBest Match:\n%s\n", result.Song)
	fmt.Printf("Similarity: %.1f%%\n", result.Similarity)
	fmt.Printf("Confidence: %s\n", result.Confidence)
}

// Run executes the full educational workflow.
func (app *App) Run(path string) (*matcher.MatchResult, error) {
	app.Logger.Log("Loading audio...")
	audioData, err := app.Reader.Read(path)
	if err != nil {
		return nil, fmt.Errorf("read audio: %w", err)
	}

	app.Logger.Log("Decoding WAV...")
	if len(audioData.Samples) == 0 {
		return nil, fmt.Errorf("audio contains no samples")
	}

	app.Logger.Log("Applying FFT...")
	fingerprint, err := app.Generator.Generate(audioData.Samples, audioData.Rate)
	if err != nil {
		return nil, fmt.Errorf("generate fingerprint: %w", err)
	}

	app.Logger.Log("Generating fingerprint...")
	app.Logger.Log("Searching...")
	result, err := app.Matcher.Match(fingerprint, app.Database.RecordsList())
	if err != nil {
		return nil, fmt.Errorf("match fingerprint: %w", err)
	}

	app.Logger.Log("Song found.")
	return result, nil
}

// FingerprintGenerator implements the educational fingerprint workflow.
type FingerprintGenerator struct{}

// Generate creates a simple fingerprint from the audio signal.
func (g *FingerprintGenerator) Generate(samples []float64, sampleRate int) (*fingerprint.Fingerprint, error) {
	if len(samples) == 0 {
		return nil, fingerprint.ErrEmptySignal
	}
	if sampleRate <= 0 {
		return nil, fmt.Errorf("invalid sample rate")
	}

	windowSize := 1024
	if len(samples) < windowSize {
		windowSize = len(samples)
	}

	peaks := make([]int, 0, 8)
	for start := 0; start+windowSize <= len(samples); start += windowSize / 2 {
		window := samples[start : start+windowSize]
		if len(window) < 16 {
			continue
		}

		// Normalize the window slightly.
		for i := range window {
			window[i] = window[i] / 32768.0
		}

		transformed := fft.FFT(window)
		magnitudes := fft.Magnitudes(transformed)

		bestIndex := 0
		for i := 1; i < len(magnitudes)/2; i++ {
			if magnitudes[i] > magnitudes[bestIndex] {
				bestIndex = i
			}
		}
		freq := int(math.Round(float64(bestIndex) * float64(sampleRate) / float64(windowSize)))
		bucketSize := 110
		bucketed := int(math.Round(float64(freq)/float64(bucketSize))) * bucketSize
		peaks = append(peaks, bucketed)
		if len(peaks) >= 8 {
			break
		}
	}

	if len(peaks) == 0 {
		return nil, fingerprint.ErrEmptySignal
	}

	return &fingerprint.Fingerprint{TopFrequencies: peaks, WindowSize: windowSize}, nil
}
