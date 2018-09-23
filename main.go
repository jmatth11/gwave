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

	// params: volume, duration, frequency
	na := notes.NewNote(7000, time.Second, notes.A)
	ng := notes.NewNote(7000, time.Second, notes.G)
	nfs := notes.NewNote(7000, time.Second, notes.Fs)
	nd := notes.NewNote(7000, time.Second, notes.D)

	sess.AddNotes(na, ng, nfs, nd)
	start := time.Now()
	sess.WriteData(pcm)
	end := time.Now()
	fmt.Printf("time: %v\n", end.Sub(start))
	err := pcm.WriteToFile("testFile.wav")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}
