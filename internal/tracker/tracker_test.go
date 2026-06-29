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
		_, err := tracker.AddItem(item)
		item.Name = "new name"
		rsl := tracker.GetItems()

		assert.NoError(t, err)
		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, rsl[0].ID, "id")
		assert.Equal(t, rsl[0].Name, "name")
	})

	t.Run("берем item из tracker, меняем имя, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		_, _ = tracker.AddItem(item)
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
		_, _ = tracker.AddItem(item)
		updErr := tracker.UpdateItem(Item{ID: "id", Name: "name2"})
		assert.NoError(t, updErr)
		rsl := tracker.GetItems()

		assert.Equal(t, len(rsl), 1)
		assert.Equal(t, rsl[0].ID, "id")
		assert.Equal(t, rsl[0].Name, "name2")
	})

	t.Run("добавляем item в tracker, удаляем item, смотрим что вернет name", func(t *testing.T) {
		t.Parallel()

		item := Item{ID: "id", Name: "name"}
		tracker := NewTracker()
		_, _ = tracker.AddItem(item)
		tracker.DeleteItem(item.ID)
		rsl := tracker.GetItems()

		assert.Equal(t, len(rsl), 0)
	})

	t.Run("error add - is exist", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		_, err := tracker.AddItem(item)
		assert.NoError(t, err)
		_, err = tracker.AddItem(item)
		assert.ErrorIs(t, err, ErrIsExist)
	})

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
}
