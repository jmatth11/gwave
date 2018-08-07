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

	na := notes.NewNote(100, time.Second, notes.A)
	ng := notes.NewNote(100, time.Second, notes.G)
	nfs := notes.NewNote(100, time.Second, notes.Fs)
	nd := notes.NewNote(100, time.Second, notes.D)

	sess.AddNotes(na, ng, nfs, nd, na, ng, nfs, nd)
	fmt.Println("added notes. length:", sess.Length())
	sess.WriteData(pcm)
	fmt.Println("finished writing data to pcm")
	err := pcm.WriteToFile("testFile.wav")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully written out file.")
}
