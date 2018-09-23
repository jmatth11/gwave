package main

import (
	"fmt"
	"time"

	"github.com/iCurlmyster/wave/format"
	"github.com/iCurlmyster/wave/notes"
	"github.com/iCurlmyster/wave/session"
)

func main() {
	pcm := format.DefaultCDPCM()
	sess := session.NewSession()

	ns := make([]*notes.Note, 24)

	// chords to Hallelujah
	// verse
	ns[0] = cChord(2)
	ns[1] = amChord(2)
	ns[2] = cChord(2)
	ns[3] = amChord(2)
	ns[4] = fChord(2)
	ns[5] = gChord(2)
	ns[6] = cChord(1)
	ns[7] = gChord(1)
	ns[8] = cChord(2)
	ns[9] = fChord(1)
	ns[10] = gChord(1)
	ns[11] = amChord(2)
	ns[12] = fChord(2)
	ns[13] = gChord(2)
	ns[14] = e7Chord(2)
	ns[15] = amChord(2)
	ns[16] = notes.SilentNote(time.Second)

	// chorus
	ns[17] = fChord(2)
	ns[18] = amChord(2)
	ns[19] = fChord(2)
	ns[20] = cChord(1)
	ns[21] = gChord(1)
	ns[22] = cChord(1)
	ns[23] = gChord(2)

	sess.AddNotes(ns...)
	fmt.Println("Data Size: ", sess.WriteData(pcm))
	if err := pcm.WriteToFile("testFile.wav"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}

func cChord(t time.Duration) *notes.Note {
	return notes.NewNote(15000, time.Second*t, notes.C, notes.E, notes.G)
}

func amChord(t time.Duration) *notes.Note {
	return notes.NewNote(15000, time.Second*t, notes.A, notes.C, notes.E)
}

func fChord(t time.Duration) *notes.Note {
	return notes.NewNote(15000, time.Second*t, notes.C, notes.F, notes.A)
}

func gChord(t time.Duration) *notes.Note {
	return notes.NewNote(15000, time.Second*t, notes.G, notes.B, notes.D, notes.G)
}

func e7Chord(t time.Duration) *notes.Note {
	return notes.NewNote(15000, time.Second*t, notes.E, notes.B, notes.D, notes.Gs)
}
