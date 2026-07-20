package audio

import "fmt"

// AudioData represents decoded PCM audio samples.
type AudioData struct {
	Samples []float64
	Rate    int
	Bits    int
	Channels int
}

// Reader loads audio from a file.
type Reader interface {
	Read(path string) (*AudioData, error)
}

// ErrUnsupportedFormat indicates that a file format is not supported.
var ErrUnsupportedFormat = fmt.Errorf("unsupported audio format")

// ErrCorruptAudio indicates the audio data could not be parsed.
var ErrCorruptAudio = fmt.Errorf("corrupt audio")
