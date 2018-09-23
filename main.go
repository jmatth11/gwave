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
	nd := notes.NewNote(17000, time.Second, notes.A)
	nfs := notes.NewNote(17000, time.Second, notes.D)
	ng := notes.NewNote(17000, time.Second, notes.G)
	na := notes.NewNote(17000, time.Second, notes.C)

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
