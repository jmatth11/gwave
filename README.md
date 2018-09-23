# gwave
Generate notes in a session and write them out to a wave file

requires:
- Go 1.10+

This project is a WIP but some of the functionality that can be done is:
- Read and Write wave files. (Currently only PCM type with a 16 byte fmt section)
- Generate notes and chords as Note objects (can change octave of Note objects as well)
- Create a Session object to store Note objects into
- Write out Note objects, stored in a Session, to a wave file.
- Supports 8bit, 16bit, and maybe 32bit (haven't verified the 32bit)
- Supports multiple channels but you normally will only need mono (1) and stereo (2)
- Supports little/big endian files (RIFF, RIFX)
- You can change the Sample rate as well, Default value is CD quality (44.1kHz)

## Run

```bash
$ go build
$ ./wave
```

## Example program

```go

package main

import (
	"fmt"
	"time"

	"github.com/iCurlmyster/wave/format"
	"github.com/iCurlmyster/wave/notes"
	"github.com/iCurlmyster/wave/session"
)

func main() {
  	// create PCM wave file
	pcm := format.DefaultCDPCM()
  	// create a new Session
  	sess := session.NewSession()
  	// create an array to hold our note objects
	ns := make([]*notes.Note, 4)

	// generate notes
  	// params: Volume, Duration, Frequencies...
	ns[0] = notes.NewNote(15000, time.Second, notes.C, notes.E, notes.G)
	ns[1] = notes.NewNote(15000, time.Second, notes.F, notes.A, notes.C)
	ns[2] = notes.NewNote(15000, time.Second, notes.C, notes.E, notes.G)
	ns[3] = notes.NewNote(15000, time.Second, notes.F, notes.A, notes.C)

  	// add all notes to session
	sess.AddNotes(ns...)
  	// write out session data to PCM wave file object
	sess.WriteData(pcm)
  	// write pcm wave file out to a new file
	if err := pcm.WriteToFile("testFile.wav"); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
```

## Project Structure

main.go has a simple example of setting up a session, assigning new notes to the session, and writing out session to a wave file.

---

Wave file objects will be in the `format` folder.

There is a `WaveWriter` interface that all wave file objects need implement to be usable by a Session. This interface is located in the `format.go` file.

---

The `notes` folder contains the Note object and constants for the frequencies of notes in the range of middle C.

--- 

The `sessions` folder contains the Session object and the methods to interact with the session.

