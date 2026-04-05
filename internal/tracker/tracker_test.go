package tracker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Add(t *testing.T) {
	t.Parallel()

	t.Run("добавляем item в tracker, меняем имя, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		tracker.AddItem(item)
		item.Name = "new name"
		rsl := tracker.GetItems()

		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, Item{ID: "id", Name: "name"}, rsl[0])
	})

	t.Run("берем item из tracker, меняем имя, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		tracker.AddItem(item)
		rsl := tracker.GetItems()
		tracker.GetItems()[0].Name = "new name"

		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, Item{ID: "id", Name: "name"}, rsl[0])
	})
}
