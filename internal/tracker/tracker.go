package tracker

import "fmt"

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("Item{id=%s, name=%s}", i.ID, i.Name)
}

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, isExist := t.indexOf(item.ID)
	if isExist {
		return Item{}, ErrIsExist
	}
	t.Items = append(t.Items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.Items[index] = item
	return nil
}

func (t *Tracker) DeleteItem(id string) {
	for i, item := range t.Items {
		if item.ID == id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}
	}
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.Items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
