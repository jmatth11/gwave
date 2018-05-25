package main

import (
	"fmt"

	"wave"
	"wave/notes"
)

var pcm = wave.DefaultCDPCM()
var oneSecond = int(pcm.Header.BytesPerSecond)
var playTime = oneSecond * 4

func main() {
	pcm.AllocateDataSize(int32(playTime))

	// i := writeOneSecondNote(notes.A, 0.5, 0, 1)
	// i = writeOneSecondNote(notes.G, 1.0, i, 2)
	// i = writeOneSecondNote(notes.Fs, 1.0, i, 3)
	// i = writeOneSecondNote(notes.D, 1.0, i, 4)
	// i = writeOneSecondNote(notes.A, 0.5, i, 5)
	// i = writeOneSecondNote(notes.G, 1.0, i, 6)
	// i = writeOneSecondNote(notes.Fs, 1.0, i, 7)
	// i = writeOneSecondNote(notes.D, 1.0, i, 8)

	i := writeChord(0, 1, notes.A, notes.G, notes.A)
	i = writeChord(i, 2, notes.D, notes.A, notes.D)
	i = writeChord(i, 3, notes.A, notes.G, notes.A)
	i = writeChord(i, 4, notes.D, notes.A, notes.D)

	err := pcm.WriteToFile("testFile.wav")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}

func writeOneSecondNote(note, octave float64, phase, times int) int {
	i := phase
	for ; i < ((oneSecond * times) - 4); i += 4 {
		pcm.SimpleStereoSingleNote(100, i, note, octave)
	}
	return i
}

// chords kind of work
func writeChord(phase, times int, note ...float64) int {
	i := phase
	for ; i < ((oneSecond * times) - 4); i += 4 {
		pcm.SimpleStereoChordNote(100, i, note...)
	}
	return i
}
