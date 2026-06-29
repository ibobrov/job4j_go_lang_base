package tracker

import (
	"bufio"
	"fmt"
	"os"
)

type Input interface {
	Get() (string, error)
}

type Output interface {
	Out(text string)
	NewLine()
	NotFoundAction()
}

type ConsoleInput struct{}

func (c ConsoleInput) Get() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", nil
	}
	return scanner.Text(), nil
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}

func (c ConsoleOutput) NewLine() {
	c.Out("")
}

func (c ConsoleOutput) NotFoundAction() {
	c.Out("not found action\n")
}
