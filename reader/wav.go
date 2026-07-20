package reader

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"sahzam/audio"
)

// WAVReader decodes PCM samples from a basic WAV file.
type WAVReader struct{}

// Read loads a WAV file and returns decoded PCM samples.
func (r *WAVReader) Read(path string) (*audio.AudioData, error) {
	if !strings.HasSuffix(strings.ToLower(path), ".wav") {
		return nil, audio.ErrUnsupportedFormat
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open audio file: %w", err)
	}
	defer file.Close()

	header := make([]byte, 44)
	if _, err := file.Read(header); err != nil {
		return nil, fmt.Errorf("read WAV header: %w", err)
	}
	if string(header[0:4]) != "RIFF" || string(header[8:12]) != "WAVE" {
		return nil, audio.ErrCorruptAudio
	}

	rate := int(binary.LittleEndian.Uint32(header[24:28]))
	bitsPerSample := int(binary.LittleEndian.Uint16(header[34:36]))
	channels := int(binary.LittleEndian.Uint16(header[22:24]))
	dataSize := int(binary.LittleEndian.Uint32(header[40:44]))

	if rate <= 0 || bitsPerSample <= 0 || channels <= 0 || dataSize <= 0 {
		return nil, audio.ErrCorruptAudio
	}

	payload := make([]byte, dataSize)
	if _, err := file.Read(payload); err != nil {
		return nil, fmt.Errorf("read WAV data: %w", err)
	}

	samples := make([]float64, 0, len(payload)/(bitsPerSample/8))
	for i := 0; i+int(bitsPerSample/8) <= len(payload); i += int(bitsPerSample / 8) {
		var sample float64
		if bitsPerSample == 8 {
			sample = float64(int(payload[i]))
		} else if bitsPerSample == 16 {
			sample = float64(int16(binary.LittleEndian.Uint16(payload[i : i+2])))
		} else {
			return nil, audio.ErrUnsupportedFormat
		}
		samples = append(samples, sample)
	}

	return &audio.AudioData{Samples: samples, Rate: rate, Bits: bitsPerSample, Channels: channels}, nil
}
