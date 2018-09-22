package main

import (
	"fmt"
	"time"

	"github.com/iCurlmyster/wave/format"
	"github.com/iCurlmyster/wave/notes"
	"github.com/iCurlmyster/wave/session"
)

var pcm = format.DefaultCDPCM()
var oneSecond = int(pcm.Header.BytesPerSecond)
var playTime = oneSecond * 4

func main() {
	pcm.AllocateDataSize(int32(playTime))

	sess := session.NewSession()

	na := notes.NewNote(5000, time.Second, notes.A)
	ng := notes.NewNote(5000, time.Second, notes.G)
	nfs := notes.NewNote(5000, time.Second, notes.Fs)
	nd := notes.NewNote(5000, time.Second, notes.D)

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
