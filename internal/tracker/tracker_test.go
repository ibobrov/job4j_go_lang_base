package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, rsl[0].ID, "id")
		assert.Equal(t, rsl[0].Name, "name")
	})

	t.Run("берем item из tracker, меняем имя, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		tracker.AddItem(item)
		rsl := tracker.GetItems()
		tracker.GetItems()[0].Name = "new name"

		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, rsl[0].ID, "id")
		assert.Equal(t, rsl[0].Name, "name")
	})

	t.Run("добавляем item в tracker, изменяем item, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		tracker.AddItem(item)
		tracker.UpdateItem(item.ID, Item{ID: "id", Name: "name2"})
		rsl := tracker.GetItems()

		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, rsl[0].ID, "id")
		assert.Equal(t, rsl[0].Name, "name2")
	})

	t.Run("добавляем item в tracker, удаляем item, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		tracker.AddItem(item)
		tracker.DeleteItem(item.ID)
		rsl := tracker.GetItems()

		assert.Equal(t, len(rsl), 0)
	})
}
