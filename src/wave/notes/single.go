package notes

import "math"

const (
	// C is middle C frequency
	C = 261.63
	// Cs is middle C sharp frequency
	Cs = 277.18
	// D is middle D frequency
	D = 293.66
	// Ds is middle D sharp frequency
	Ds = 311.13
	// E is middle E frequency
	E = 329.63
	// F is middle F frequency
	F = 349.23
	// Fs is middle F sharp frequency
	Fs = 369.99
	// G is middle G frequency
	G = 392.00
	// Gs is middle G sharp frequency
	Gs = 415.30
	// A is middle A frequency
	A = 440.00
	// As is middle A sharp frequency
	As = 466.16
	// B is middle B frequency
	B = 493.88
)

// CreateNote creates a note at a certain octave from middle C at a certain time point.
func CreateNote(t int, octave, bps, freq float64) float64 {
	// times bps by <1 or >=1 to lower or raise octave respectively
	i := 2.0 * math.Pi * freq / (bps * octave)
	return math.Sin(i * float64(t)) //+ 1.0
}
