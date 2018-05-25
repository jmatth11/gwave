package notes

import (
	"math"
	"wave"
)

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

// Note wrapper around things needed for creating sine wave for a note
type Note struct {
	Volume    int16
	Frequency float64
	Octave    float64
}

// CreateNote creates a note at a certain octave from middle C at a certain time point.
func CreateNote(t int, octave, bps, freq float64) float64 {
	// times bps by <1 or >=1 to lower or raise octave respectively
	i := 2.0 * math.Pi * freq / (bps * octave)
	return math.Sin(i * float64(t)) //+ 1.0
}

// AtTime grabs sin wave at time t
func (note Note) AtTime(t int, bps float64) float64 {
	return NoteAtTime(t, bps, note)
}

// NoteAtTime grabs sin wave at time t
func NoteAtTime(t int, bps float64, note Note) float64 {
	// times bps by <1 or >=1 to lower or raise octave respectively
	i := 2.0 * math.Pi * note.Frequency / (bps * note.Octave)
	return math.Sin(i * float64(t))
}

// ToData works for 16 bit sample note. same note for left and right side
// Example Volume and Frequency could be 32000 and 440.0 respectively
func (note Note) ToData(vol int16, index int, header wave.Header) []byte {
	val := int16(float64(vol) * note.AtTime(index, float64(header.BytesPerSecond)))
	return Int16ToBytes(val)
}
