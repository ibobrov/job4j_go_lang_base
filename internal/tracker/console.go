package tracker

import (
	"bufio"
	"fmt"
	"os"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
	OutActions(actions map[string]UseCase)
	NewLine()
	NotFoundAction()
}

type ConsoleInput struct{}

func (c ConsoleInput) Get() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// block IO
	}
	return scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}

func (c ConsoleOutput) OutActions(actions map[string]UseCase) {
	c.Out("---------------")
	c.Out("Select action\n")
	for i := range actions {
		c.Out(i)
	}
	c.Out("exit\n")
}

func (c ConsoleOutput) NewLine() {
	c.Out("")
}

func (c ConsoleOutput) NotFoundAction() {
	c.Out("not found action\n")
}
