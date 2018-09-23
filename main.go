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

	ns := make([]*notes.Note, 4)

	// params: volume, duration, frequency
	ns[0] = notes.NewNote(15000, time.Second, notes.F, notes.A, notes.B)
	ns[1] = notes.NewNoteWithOctave(15000, 0.5, time.Second, notes.F)
	ns[2] = notes.NewNote(15000, time.Second, notes.D, notes.A, notes.G)
	ns[3] = notes.NewNoteWithOctave(15000, 0.5, time.Second, notes.D)

	sess.AddNotes(ns...)
	fmt.Println("Data Size: ", sess.WriteData(pcm))
	if err := pcm.WriteToFile("testFile.wav"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}
