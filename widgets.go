package main

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func AddTodoWidget() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter todo title..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20

	return ti
}
