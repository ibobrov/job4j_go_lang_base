package tracker

import "github.com/google/uuid"

type UseCase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUseCase struct{}

func (u AddUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	item, err := tracker.AddItem(Item{Name: name, ID: id})
	if err == nil {
		out.Out("item was added: " + item.toString())
	} else {
		out.Out(err.Error())
	}
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id:")
	id := in.Get()

	out.Out("enter new name:")
	name := in.Get()
	err := tracker.UpdateItem(Item{Name: name, ID: id})
	if err == nil {
		out.Out("item was updated")
	} else {
		out.Out(err.Error())
	}
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id:")
	id := in.Get()
	tracker.DeleteItem(id)
}
