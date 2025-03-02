package main

import (
	"fmt"

	"example.org/todo-term/todo"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos      []todo.Todo
	cursor     int
	todoText   textinput.Model
	addingTodo bool
}

func initialModel() model {

	return model{
		todos:      todo.InitialTodos,
		addingTodo: false,
		todoText:   AddTodoWidget(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			handleUp(&m)
		case "down", "s":
			handleDown(&m)
		case "a":
			m.addingTodo = true
		case "enter":
			handleEnter(&m)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if m.addingTodo {
		m.todoText, cmd = m.todoText.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	s := "Welcome to term-todo\n\n"
	doneTodos := []todo.Todo{}

	for index, todo := range m.todos {

		if todo.Done {
			doneTodos = append(doneTodos, todo)
			continue
		}
		current := " "
		if index == m.cursor {
			current = "x"
		}

		s += fmt.Sprintf("[%s] %s\n", current, todo.Title)
	}

	s += "\n-----Done----\n"

	for _, todo := range doneTodos {

		s += fmt.Sprintf("[✅] %s\n", todo.Title)
	}

	s += "\n"

	if m.addingTodo {
		s += m.todoText.View()
	}

	return s
}
