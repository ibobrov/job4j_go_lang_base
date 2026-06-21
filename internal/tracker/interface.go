package tracker

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) OutActions(actions map[string]UseCase) {
	u.Out.Out("---------------")
	u.Out.Out("Select action\n")
	for i := range actions {
		u.Out.Out(i)
	}
	u.Out.Out("exit\n")
}

func (u UI) Run() {
	actions := map[string]UseCase{
		"add":    AddUseCase{},
		"get":    GetUseCase{},
		"update": UpdateUseCase{},
		"delete": DeleteUseCase{},
	}

	for {
		u.OutActions(actions)
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.NotFoundAction()
			continue
		}

		action.Done(u.In, u.Out, u.Tracker)
		u.Out.NewLine()
	}
}
