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
	fmt.Println("Data Size: ", sess.WriteData(pcm))
	if err := pcm.WriteToFile("testFile.wav"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}
