// A Chord Interval is how the members of the chord are counted, from 1 (the "root") to e.g. 3 (the "third") or 5 (the "fifth")
package chord

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devthane/music-theory/note"
)

func TestInterval(t *testing.T) {
	assert.Equal(t, 16, len(intervalOrder))
}

func TestForAllIn(t *testing.T) {
	tones := map[Interval]note.Class{
		I2: note.Ds,
		I5: note.G,
		I7: note.As,
	}
	forAllIn(tones, func(class note.Class) {
		assert.NotEmpty(t, class)
	})
}
