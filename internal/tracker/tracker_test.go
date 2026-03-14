package tracker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"
		assert.Equal(t,
			[]Item{item},
			tracker.GetItems(),
		)
	})
}
