package session

import (
	"errors"
	"sync"
	"time"

	"github.com/iCurlmyster/wave/format"
	"github.com/iCurlmyster/wave/notes"
)

// Session represents a collection of notes to be played in sequence
type Session struct {
	length         time.Duration
	phase          int
	noteCollection []*notes.Note
}

// NewSession creates a new session object
func NewSession() *Session {
	return &Session{
		length:         0,
		phase:          0,
		noteCollection: make([]*notes.Note, 0),
	}
}

// Length returns the current duration of the session
func (sess *Session) Length() time.Duration {
	return sess.length
}

// AddNotes handles adding new Notes to the session
func (sess *Session) AddNotes(ns ...*notes.Note) {
	dur := sess.length
	for _, n := range ns {
		if n.Length < 0 {
			panic(errors.New("cannot add a note with a negative duration"))
		}
		dur = dur + n.Length
	}
	sess.length = dur
	sess.noteCollection = append(sess.noteCollection, ns...)
}

// WriteData writes session data out to wave writer object
func (sess *Session) WriteData(wWriter format.WaveWriter) int32 {
	bps := wWriter.FileHeader().BytesPerSecond
	size := int32((sess.length / time.Second) * time.Duration(bps))
	wWriter.AllocateDataSize(size)
	dur := 0
	wg := sync.WaitGroup{}
	for i := 0; i < len(sess.noteCollection); i++ {
		dur += addNoteParallel(dur, sess.noteCollection[i], &wg, wWriter)
	}
	wg.Wait()
	return size
}
