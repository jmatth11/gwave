package notes

import (
	"github.com/iCurlmyster/wave/format"
)

// ChordAtTime generates a chord sine wave of the notes given
func ChordAtTime(index int, pcm *format.PCM, ns ...Note) float64 {
	sum := 0.0
	vol := int16(0)
	for _, n := range ns {
		if n.Volume > vol {
			vol = n.Volume
		}
		sum += n.AtTime(index, float64(pcm.BytesPerSecond))
	}
	val := float64(vol/int16(len(ns))) * sum
	return val
}
