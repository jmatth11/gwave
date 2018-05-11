package main

import (
	"fmt"

	"./src/wave"
	"./src/wave/notes"
)

var oneSecond = wave.CdSampleRate * 4
var playTime = oneSecond * 8
var pcm = wave.DefaultCDPCM()

func main() {
	pcm.AllocateDataSize(int32(playTime))

	i := writeOneSecondNote(notes.A, 1.0, 0, 1)
	i = writeOneSecondNote(notes.G, 2.0, i, 2)
	i = writeOneSecondNote(notes.Fs, 2.0, i, 3)
	i = writeOneSecondNote(notes.D, 2.0, i, 4)
	i = writeOneSecondNote(notes.A, 1.0, i, 5)
	i = writeOneSecondNote(notes.G, 2.0, i, 6)
	i = writeOneSecondNote(notes.Fs, 2.0, i, 7)
	i = writeOneSecondNote(notes.D, 2.0, i, 8)

	err := pcm.WriteToFile("testFile.wav")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}

func writeOneSecondNote(note, octave float64, phase, times int) int {
	i := phase
	for ; i < ((oneSecond * times) - 4); i += 4 {
		pcm.SimpleStereoSingleNote(500, i, note, octave)
	}
	return i
}
