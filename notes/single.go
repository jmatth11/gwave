package notes

import (
	"math"
	"time"
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
	Volume    float64
	Frequency []float64
	Length    time.Duration
}

// SilentNote generates a silent note of defined length
func SilentNote(length time.Duration) *Note {
	return &Note{
		Volume:    0.0,
		Frequency: []float64{0.0},
		Length:    length,
	}
}

// NewNote generates a new note object
func NewNote(vol float64, len time.Duration, freq ...float64) *Note {
	return &Note{
		Volume:    vol,
		Frequency: freq,
		Length:    len,
	}
}

// AtTime grabs sin value at time t
func (note Note) AtTime(t int) float64 {
	return NoteAtTime(t, note)
}

// NoteAtTime grabs sin value at time t
func NoteAtTime(t int, note Note) float64 {
	sum := 0.0
	for i := 0; i < len(note.Frequency); i++ {
		sum += 2.0 * math.Pi / note.Frequency[i]
	}
	return math.Sin(sum * float64(t))
}

// ToData generates a float64 value representation of the note
// Accounts for mutliple notes by dividing the volume by the number of notes
func (note Note) ToData(index int) float64 {
	freqLen := len(note.Frequency)
	vol := note.Volume
	if freqLen > 1 {
		vol = vol / float64(freqLen)
	}
	return vol * note.AtTime(index)
}
