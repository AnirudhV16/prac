package types

import "fmt"

/*
Build a smart home controller:

Command interface with Execute() and Undo()
Light receiver with TurnOn() and TurnOff()
TurnOnCommand and TurnOffCommand — both implement Command
RemoteControl invoker with PressButton(cmd Command) and PressUndo()
In main — turn light on, turn light off, undo twice — show the light goes back on
*/

type Command interface {
	Execute()
	Undo()
}

//receiver (actual work doer)
type Light struct {
	state int // 0 is OFF and 1 is ON
}

func (l *Light) GetState() int {
	return l.state
}

func NewLight() *Light {
	return &Light{state: 0}
}
func (l *Light) TurnOn() {
	if l.state == 0 {
		l.state = 1
		fmt.Println("light turned on....")
		return
	}
	fmt.Println("light already turned on....")
}

func (l *Light) TurnOff() {
	if l.state == 1 {
		l.state = 0
		fmt.Println("light turned off....")
		return
	}
	fmt.Println("light already turned off....")
}

//concrete commands
type TurnOnCommand struct {
	Light *Light
}

func (t *TurnOnCommand) Execute() {
	t.Light.TurnOn()
	t.Light.GetState()
}

func (t *TurnOnCommand) Undo() {
	v := t.Light.GetState()
	if v == 1 {
		t.Light.state = 0
		fmt.Println("on state undone....")
	}
}

type TurnOffCommand struct {
	Light *Light
}

func (t *TurnOffCommand) Execute() {
	t.Light.TurnOff()
	t.Light.GetState()
}

func (t *TurnOffCommand) Undo() {
	v := t.Light.GetState()
	if v == 0 {
		t.Light.state = 1
		fmt.Println("off state undone....")
	}
}

type RemoteControl struct {
	History []Command
}

func (r *RemoteControl) PressButton(cmd Command) {
	cmd.Execute()
	r.History = append(r.History, cmd)
}

func (r *RemoteControl) PressUndo() {
	if len(r.History) == 0 {
		fmt.Println("nothing to undo")
		return
	}
	last := r.History[len(r.History)-1]
	last.Undo()
	r.History = r.History[:len(r.History)-1]
}
