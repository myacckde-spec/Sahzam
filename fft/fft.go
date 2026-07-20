package fft

import "math"

// FFT computes a simple forward FFT over complex-valued input.
// This implementation is intentionally small and educational.
func FFT(values []float64) []complex128 {
	if len(values) == 0 {
		return nil
	}

	n := len(values)
	if n == 1 {
		return []complex128{complex(values[0], 0)}
	}

	// Cooley-Tukey radix-2 decomposition.
	if n%2 != 0 {
		// For the MVP, pad with zeros when the window size is odd.
		padded := make([]float64, n+1)
		copy(padded, values)
		values = padded
		n = len(values)
	}

	even := make([]float64, n/2)
	odd := make([]float64, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = values[2*i]
		odd[i] = values[2*i+1]
	}

	evenFFT := FFT(even)
	oddFFT := FFT(odd)

	result := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		angle := -2 * math.Pi * float64(k) / float64(n)
		w := complex(math.Cos(angle), math.Sin(angle))
		term := oddFFT[k] * w
		result[k] = evenFFT[k] + term
		result[k+n/2] = evenFFT[k] - term
	}
	return result
}

// Magnitudes returns the magnitudes of the FFT output.
func Magnitudes(values []complex128) []float64 {
	magnitudes := make([]float64, len(values))
	for i, v := range values {
		magnitudes[i] = real(v)*real(v) + imag(v)*imag(v)
	}
	return magnitudes
}
