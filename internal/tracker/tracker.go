package tracker

import "fmt"

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) UpdateItem(id string, new Item) {
	for i, item := range t.Items {
		if item.ID == id {
			t.Items[i] = new
		}
	}
}

func (t *Tracker) DeleteItem(id string) {
	for i, item := range t.Items {
		if item.ID == id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}
	}
}
